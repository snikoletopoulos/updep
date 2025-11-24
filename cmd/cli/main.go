package main

import (
	"npmupdate/pkg/components/row"
	"npmupdate/pkg/config"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type AppModel struct {
	spinner spinner.Model
	help    help.Model
	rows    []row.Row
	cursor  int
}

func NewAppModel() AppModel {
	s := spinner.New()
	s.Spinner = spinner.Points
	s.Style = lipgloss.NewStyle().Foreground(config.Theme.Mauve)

	return AppModel{
		spinner: s,
		help:    help.New(),
		rows:    []row.Row{},
		cursor:  0,
	}
}

func (m AppModel) Init() tea.Cmd {
	return tea.Batch(m.spinner.Tick, getOutdatedPackages)
}

func (m AppModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case OutdatedPackagesMsg:
		packages := msg

		columnWidth := []int{0, 0, 0, 0}
		for _, p := range packages {
			columnWidth[0] = max(columnWidth[0], lipgloss.Width(p.Name))
			columnWidth[1] = max(columnWidth[1], lipgloss.Width(p.Wanted.String()))
			columnWidth[2] = max(columnWidth[2], lipgloss.Width(p.Latest.String()))
			columnWidth[3] = max(columnWidth[3], lipgloss.Width(p.Current.String()))
		}

		for _, p := range packages {
			m.rows = append(m.rows, row.New(p, columnWidth))
		}
	case tea.KeyMsg:
		cmds = append(cmds, m.handleKeyPress(msg))
	}

	if m.cursor < len(m.rows)-1 {
		rowModel, cmd := m.rows[m.cursor].Update(msg)
		m.rows[m.cursor] = rowModel.(row.Row)
		cmds = append(cmds, cmd)
	}

	var cmd tea.Cmd
	m.spinner, cmd = m.spinner.Update(msg)
	cmds = append(cmds, cmd)

	return m, tea.Batch(cmds...)
}

func (m AppModel) View() string {
	if len(m.rows) == 0 {
		return m.spinner.View() + " Getting outdated packages"
	}

	renderRows := []string{}
	for i, p := range m.rows {
		if i != m.cursor {
			renderRows = append(renderRows, p.View())
			continue
		}

		renderRows = append(
			renderRows,
			row.ActiveRowStyle.Render(p.View()),
		)
	}

	return lipgloss.JoinVertical(
		lipgloss.Left,
		lipgloss.JoinVertical(lipgloss.Left, renderRows...),
		m.help.View(keyMap),
	)
}

func main() {
	p := tea.NewProgram(NewAppModel())
	if _, err := p.Run(); err != nil {
		panic(err)
	}
}
