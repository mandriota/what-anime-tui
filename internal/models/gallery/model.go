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
	"strings"
	"time"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/paginator"
	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	"github.com/mandriota/what-anime-tui/internal/ascii"
	"github.com/mandriota/what-anime-tui/internal/config"
	"github.com/mandriota/what-anime-tui/internal/fetcher"
)

type searchFinishedMsg struct {
	err error
}

type style struct {
	base       lipgloss.Style
	card       lipgloss.Style
	gallery    lipgloss.Style
	deck       lipgloss.Style
	statePanel lipgloss.Style
}

func newStyle(cfg config.AppearanceConfig) style {
	base := lipgloss.NewStyle().
		Background(cfg.Background).
		Foreground(cfg.Foreground)

	return style{
		base: base,
		card: lipgloss.NewStyle().
			PaddingLeft(4),
		gallery: lipgloss.NewStyle().
			Padding(1, 0, 1, 0),
		deck: base.Copy().
			Height(10).
			BorderTopForeground(cfg.Border.Foreground).
			BorderStyle(lipgloss.ThickBorder()).
			BorderTop(true).
			Bold(true),
		statePanel: base.Copy().
			Align(lipgloss.Center),
	}
}

type Model struct {
	fetcher  fetcher.Fetcher
	response *fetcher.Response

	KeyMap KeyMap

	textInput textinput.Model
	paginator paginator.Model
	spinner   spinner.Model
	help      help.Model

	altScreen bool
	searching bool
	cliQuit   bool

	err error

	style style

	cfg config.GeneralConfig
}

func New(cfg config.GeneralConfig, cliQuit bool, path string) Model {
	am := Model{
		fetcher:  fetcher.New(cfg.Fetcher),
		response: new(fetcher.Response),
		KeyMap:   DefaultKeyMap,
		spinner: spinner.New(spinner.WithSpinner(spinner.Spinner{
			Frames: ascii.ArtTelescope,
			FPS:    time.Millisecond * 1500 / time.Duration(len(ascii.ArtTelescope)),
		})),
		help:    help.New(),
		style:   newStyle(cfg.Appearance),
		cfg:     cfg,
		cliQuit: cliQuit,
	}

	am.textInput = textinput.New()
	am.textInput.Placeholder = "Enter File or URL to Search"
	am.textInput.Prompt = ""
	am.textInput.SetValue(path)
	am.textInput.Focus()

	am.paginator = paginator.New()
	am.paginator.Type = paginator.Dots
	am.paginator.KeyMap = am.KeyMap.Gallery.Paginator
	am.paginator.InactiveDot = am.style.base.Copy().
		Faint(true).
		Render("•")
	am.paginator.ActiveDot = am.style.base.Copy().
		Bold(true).
		Render("•")

	return am
}

func (m Model) Init() tea.Cmd {
	return func() tea.Msg {
		if path := m.textInput.Value(); path != "" {
			return tea.KeyMsg(tea.Key{
				Type: tea.KeyEnter,
			})
		}
		return nil
	}
}

func (m Model) Update(msg tea.Msg) (_ tea.Model, cmd tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.KeyMap.Quit):
			return m, tea.Quit
		case key.Matches(msg, m.KeyMap.Search):
			m.textInput.Blur()
			m.response.Result = m.response.Result[:0]
			m.err = nil

			path := m.textInput.Value()
			if path == "" || m.searching {
				return m, nil
			}

			cmd = tea.Batch(
				m.spinner.Tick,
				func() tea.Msg {
					switch {
					case strings.HasPrefix(path, "http"):
						m.err = m.fetcher.FetchByURL(m.response, path)
					default:
						m.err = m.fetcher.FetchByFile(m.response, path)
					}

					return searchFinishedMsg{err: m.err}
				},
			)

			if m.cliQuit {
				cmd = tea.Sequence(cmd, tea.Quit)
			}

			m.searching = true
			return m, cmd
		case key.Matches(msg, m.KeyMap.Help):
			m.help.ShowAll = !m.help.ShowAll
		case key.Matches(msg, m.KeyMap.AltScreen):
			m.altScreen = !m.altScreen
			if m.altScreen {
				return m, tea.EnterAltScreen
			}
			return m, tea.ExitAltScreen
		}

		if m.textInput.Focused() {
			switch {
			case key.Matches(msg, m.KeyMap.Form.Blur):
				m.textInput.Blur()
				return m, nil
			}
			m.textInput, cmd = m.textInput.Update(msg)
			return m, cmd
		}

		switch {
		case key.Matches(msg, m.KeyMap.Gallery.Blur):
			m.textInput.Focus()
		case key.Matches(msg, m.KeyMap.Gallery.Paginator.NextPage, m.KeyMap.Gallery.Paginator.PrevPage):
			m.paginator, cmd = m.paginator.Update(msg)
			return m, cmd
		}
	case tea.WindowSizeMsg:
		m.style.deck.Width(msg.Width)
		m.style.statePanel.Width(msg.Width)
		m.help.Width = msg.Width
		m.textInput.Width = msg.Width - 1
		return m, nil
	case searchFinishedMsg:
		m.err = msg.err
		m.paginator.SetTotalPages(len(m.response.Result))
		m.paginator.Page = 0
		m.searching = false
	case spinner.TickMsg:
		if !m.searching {
			return m, nil
		}
		m.spinner, cmd = m.spinner.Update(msg)
		return m, cmd
	}

	return m, nil
}

func (m Model) View() (deck string) {
	switch {
	case m.err != nil:
		deck = m.style.statePanel.Render(m.err.Error())
	case m.searching:
		deck = m.style.statePanel.Render("SEARCHING\n" + m.spinner.View())
	case len(m.response.Result) == 0:
		deck = m.style.statePanel.Render("NO RESULTS\n" + ascii.ArtNoResults)
	default:
		deck = m.style.statePanel.Render(m.paginator.View()) + "\n" +
			m.style.card.Render(m.response.Result[m.paginator.Page].View()+"\n")
	}

	return m.style.gallery.Render(m.textInput.View()+"\n"+
		m.style.deck.Render(deck),
	) + "\n" +
		m.help.View(m.KeyMap)
}
