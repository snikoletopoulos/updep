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
	loading      bool
	spinner      spinner.Model
	help         help.Model
	rows         []row.Row
	cursor       int
	columnWidths [config.ColumnCount]int
}

func NewAppModel() AppModel {
	s := spinner.New()
	s.Spinner = spinner.Points
	s.Style = lipgloss.NewStyle().Foreground(config.Theme.Mauve)

	return AppModel{
		loading:      true,
		spinner:      s,
		help:         help.New(),
		rows:         []row.Row{},
		columnWidths: [config.ColumnCount]int{},
		cursor:       0,
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
		m.columnWidths = row.CalculateColumnWidths(packages)
		m.rows = make([]row.Row, len(packages))
		for i, p := range packages {
			m.rows[i] = row.New(p, m.columnWidths)
		}
		m.loading = false
	case tea.KeyMsg:
		cmds = append(cmds, m.handleKeyPress(msg))
	}

	if m.cursor < len(m.rows) {
		rowModel, cmd := m.rows[m.cursor].Update(msg)
		m.rows[m.cursor] = rowModel.(row.Row)
		cmds = append(cmds, cmd)
	}

	if m.loading {
		var cmd tea.Cmd
		m.spinner, cmd = m.spinner.Update(msg)
		cmds = append(cmds, cmd)
	}

	return m, tea.Batch(cmds...)
}

func (m AppModel) View() string {
	if m.loading {
		return m.spinner.View() + " Getting outdated packages"
	}

	renderRows := make([]string, len(m.rows))
	for i, p := range m.rows {
		renderRows[i] = p.View()
		if i == m.cursor {
			renderRows[i] = row.ActiveRowStyle.Render(p.View())
		}
	}

	return lipgloss.JoinVertical(
		lipgloss.Left,
		lipgloss.JoinVertical(lipgloss.Left, renderRows...),
		m.help.View(keyMap),
	)
}

func main() {
	p := tea.NewProgram(NewAppModel())
	defer p.Kill()
	if _, err := p.Run(); err != nil {
		panic(err)
	}
}
