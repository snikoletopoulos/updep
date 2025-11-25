package row

import (
	"npmupdate/pkg/config"
	packagemodel "npmupdate/pkg/models/package"
	"npmupdate/pkg/models/version"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Row struct {
	pkg          packagemodel.Package
	target       *version.Version
	ColumnWidths [config.ColumnCount]int
}

func New(pkg packagemodel.Package, columnWidths [config.ColumnCount]int) Row {
	return Row{
		pkg:          pkg,
		ColumnWidths: columnWidths,
	}
}

func (r Row) Init() tea.Cmd {
	return nil
}

func (r Row) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		cmds = append(cmds, r.handleKeyPress(msg))
	}

	return r, tea.Batch(cmds...)
}

func (r Row) View() string {
	var cells []string
	for index, cell := range r.getCellStyles() {
		width := r.ColumnWidths[index] + config.ColumnGap
		if index == len(r.ColumnWidths)-1 {
			width -= config.ColumnGap
		}
		cells = append(cells, lipgloss.PlaceHorizontal(width, lipgloss.Left, cell))
	}

	return lipgloss.JoinHorizontal(lipgloss.Center, cells...)
}

func (r Row) getCellStyles() [config.ColumnCount]string {
	var nameCellStyle, wantedCellStyle, latestCellStyle lipgloss.Style

	if r.pkg.Current.Compare(r.pkg.Wanted) == -1 {
		nameCellStyle = needUpdateStyle
	} else if r.pkg.Current.Compare(r.pkg.Wanted) == 1 {
		nameCellStyle = errorVersionStyle
	} else {
		nameCellStyle = optionalUpdateStyle
	}

	if r.target != nil {
		switch *r.target {
		case r.pkg.Wanted:
			wantedCellStyle = wantedStyle
		case r.pkg.Latest:
			latestCellStyle = latestStyle
		}
	}

	return [config.ColumnCount]string{
		nameCellStyle.Render(r.pkg.Name),
		wantedCellStyle.Render(r.pkg.Wanted.String()),
		latestCellStyle.Render(r.pkg.Latest.String()),
		r.pkg.Current.String(),
	}
}
