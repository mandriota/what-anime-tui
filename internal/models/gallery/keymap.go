// Copyright 2023 Mark Mandriota
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package gallery

import (
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/paginator"
)

type KeyMapForm struct {
	Blur key.Binding
}

type KeyMapGallery struct {
	Paginator paginator.KeyMap
	Blur      key.Binding
}

type KeyMap struct {
	Form      KeyMapForm
	Gallery   KeyMapGallery
	AltScreen key.Binding
	Search    key.Binding
	Help      key.Binding
	Quit      key.Binding
}

func (k KeyMap) ShortHelp() []key.Binding {
	return []key.Binding{k.Help, k.Quit}
}

func (k KeyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.Help, k.Quit, k.Search, k.AltScreen},
		{k.Gallery.Blur, k.Form.Blur, k.Gallery.Paginator.PrevPage, k.Gallery.Paginator.NextPage},
	}
}

var DefaultKeyMap = KeyMap{
	Form: KeyMapForm{
		Blur: key.NewBinding(
			key.WithKeys("down", "esc"),
			key.WithHelp("↓/esc", "Form Blur"),
		),
	},
	Gallery: KeyMapGallery{
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
		Blur: key.NewBinding(
			key.WithKeys("up", "j"),
			key.WithHelp("↑/j", "Form Focus"),
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
	Help: key.NewBinding(
		key.WithKeys("ctrl+g"),
		key.WithHelp("ctrl+g", "Toggle Help"),
	),
	Quit: key.NewBinding(
		key.WithKeys("ctrl+c"),
		key.WithHelp("ctrl+c", "Quit"),
	),
}
