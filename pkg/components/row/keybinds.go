package row

import (
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
)

type KeyMap struct {
	Wanted, Latest, ToggleSelect key.Binding
}

var keyMap = KeyMap{
	Wanted: key.NewBinding(
		key.WithKeys("w"),
		key.WithHelp("w", "update to wanted version"),
	),
	Latest: key.NewBinding(
		key.WithKeys("l"),
		key.WithHelp("l", "update to latest version"),
	),
	ToggleSelect: key.NewBinding(
		key.WithKeys(" "),
		key.WithHelp("space", "toggle"),
	),
}

func (r *Row) handleKeyPress(msg tea.KeyMsg) tea.Cmd {
	switch {
	case key.Matches(msg, keyMap.Wanted):
		r.target = &r.pkg.Wanted
	case key.Matches(msg, keyMap.Latest):
		r.target = &r.pkg.Latest
	case key.Matches(msg, keyMap.ToggleSelect):
		if r.target != nil {
			r.target = nil
			break
		}
		if r.pkg.Current.Compare(r.pkg.Wanted) >= 0 {
			r.target = &r.pkg.Latest
		} else {
			r.target = &r.pkg.Wanted
		}
	}
	return nil
}
