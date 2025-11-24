package row

import (
	"npmupdate/pkg/entities"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var gap int = 4

type Row struct {
	pkg          entities.Package
	columnWidths []int
	target       *entities.Version
}

func New(pkg entities.Package, columnWidths []int) Row {
	return Row{
		pkg:          pkg,
		columnWidths: columnWidths,
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
	nameCellStyle := lipgloss.NewStyle()
	wantedCellStyle := lipgloss.NewStyle()
	latestCellStyle := lipgloss.NewStyle()

	if r.target != nil {
		switch *r.target {
		case r.pkg.Wanted:
			wantedCellStyle = wantedStyle
		case r.pkg.Latest:
			latestCellStyle = latestStyle
		}
	}

	if r.pkg.Current.Compare(r.pkg.Wanted) == -1 {
		nameCellStyle = needUpdateStyle
	} else if r.pkg.Current.Compare(r.pkg.Wanted) == 1 {
		nameCellStyle = errorVersionStyle
	} else {
		nameCellStyle = optionalUpdateStyle
	}

	nameCell := lipgloss.PlaceHorizontal(
		r.columnWidths[0]+gap,
		lipgloss.Left,
		nameCellStyle.Render(r.pkg.Name),
	)
	wantedCell := lipgloss.PlaceHorizontal(
		r.columnWidths[1]+gap,
		lipgloss.Left,
		wantedCellStyle.Render(r.pkg.Wanted.String()),
	)
	latestCell := lipgloss.PlaceHorizontal(
		r.columnWidths[2]+gap,
		lipgloss.Left,
		latestCellStyle.Render(r.pkg.Latest.String()),
	)
	currentCell := lipgloss.PlaceHorizontal(
		r.columnWidths[3],
		lipgloss.Left,
		r.pkg.Current.String(),
	)

	return lipgloss.JoinHorizontal(
		lipgloss.Center,
		nameCell,
		wantedCell,
		latestCell,
		currentCell,
	)
}
