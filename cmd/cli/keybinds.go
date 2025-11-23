package main

import (
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
)

type KeyMap struct {
	Up, Down, Quit key.Binding
}

var DefaultKeyMap = KeyMap{
	Up: key.NewBinding(
		key.WithKeys("k", "up"),
		key.WithHelp("↑/k", "move up"),
	),
	Down: key.NewBinding(
		key.WithKeys("j", "down"),
		key.WithHelp("↓/j", "move down"),
	),
	Quit: key.NewBinding(
		key.WithKeys("q", "ctrl+c", "esc"),
		key.WithHelp("↓/j", "move down"),
	),
}

func (m *AppModel) handleKeyPress(msg tea.KeyMsg) tea.Cmd {
	switch {
	case key.Matches(msg, DefaultKeyMap.Up):
		if m.cursor > 0 {
			m.cursor -= 1
		}
	case key.Matches(msg, DefaultKeyMap.Down):
		if m.cursor < len(m.rows)-1 {
			m.cursor += 1
		}
	case key.Matches(msg, DefaultKeyMap.Quit):
		return tea.Quit
	}
	return nil
}
