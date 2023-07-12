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
	Blur      key.Binding
	Help      key.Binding
	Quit      key.Binding
}

func (k KeyMap) ShortHelp() []key.Binding {
	return []key.Binding{k.Help, k.Quit}
}

func (k KeyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.Help, k.Quit, k.Search, k.AltScreen},
		{k.Focus, k.Blur, k.Paginator.PrevPage, k.Paginator.NextPage},
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
		key.WithKeys("ctrl+s"),
		key.WithHelp("ctrl+s", "Toggle AltScreen"),
	),
	Search: key.NewBinding(
		key.WithKeys("enter"),
		key.WithHelp("enter", "Search"),
	),
	Focus: key.NewBinding(
		key.WithKeys("j"),
		key.WithHelp("j", "Form Focus"),
	),
	Blur: key.NewBinding(
		key.WithKeys("esc"),
		key.WithHelp("esc", "Form Blur"),
	),
	Help: key.NewBinding(
		key.WithKeys("ctrl+g"),
		key.WithHelp("ctrl+g", "Toggle Help"),
	),
	Quit: key.NewBinding(
		key.WithKeys("ctrl+c", "ctrl+q"),
		key.WithHelp("ctrl+c/ctrl+q", "Quit"),
	),
}
