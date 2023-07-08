package anideck

import (
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/paginator"
)

type KeyMap struct {
	Paginator paginator.KeyMap
	AltScreen key.Binding
	Search    key.Binding
	Focus     key.Binding
	Help      key.Binding
	Quit      key.Binding
}

func (k KeyMap) ShortHelp() []key.Binding {
	return []key.Binding{k.Help, k.Quit}
}

func (k KeyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.Help, k.Quit, k.Focus, k.Search},
		{k.AltScreen, k.Paginator.PrevPage, k.Paginator.NextPage},
	}
}

var DefaultKeyMap = KeyMap{
	Paginator: paginator.KeyMap{
		NextPage: key.NewBinding(
			key.WithKeys("right", "l"),
			key.WithHelp("→/l", "Next Card"),
		),
		PrevPage: key.NewBinding(
			key.WithKeys("left", "h"),
			key.WithHelp("←/h", "Prev Card"),
		),
	},
	AltScreen: key.NewBinding(
		key.WithKeys(" "),
		key.WithHelp("space", "Toggle AltScreen"),
	),
	Search: key.NewBinding(
		key.WithKeys("enter"),
		key.WithHelp("enter", "Search"),
	),
	Focus: key.NewBinding(
		key.WithKeys("tab"),
		key.WithHelp("tab", "Switch Focus"),
	),
	Help: key.NewBinding(
		key.WithKeys("?"),
		key.WithHelp("?", "Toggle Help"),
	),
	Quit: key.NewBinding(
		key.WithKeys("ctrl+c", "ctrl+q", "esc"),
		key.WithHelp("ctrl+c/ctrl+q/esc", "Quit"),
	),
}
