package anideck

import (
	"strings"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/paginator"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/mandriota/what-anime-tui/internal/fetcher"
	"github.com/muesli/termenv"
)

const (
	APIBaseURL        = `https://api.trace.moe/search?anilistInfo`
	APIByURLParameter = `&url=`
)

var (
	styleBase = lipgloss.NewStyle().
			Background(lipgloss.ANSIColor(termenv.ANSICyan)).
			Foreground(lipgloss.ANSIColor(termenv.ANSIBrightWhite))
	styleCard = lipgloss.NewStyle().
			PaddingLeft(4)
)

type searchFinishedMsg struct{}

type Model struct {
	fetcher  fetcher.Fetcher
	response *fetcher.Response

	KeyMap KeyMap

	textInput textinput.Model
	paginator paginator.Model
	help      help.Model

	altScreen bool
	searching bool

	styleWidget     lipgloss.Style
	styleStatePanel lipgloss.Style
}

func New(path string) Model {
	am := Model{
		fetcher:   fetcher.New(),
		response:  new(fetcher.Response),
		KeyMap:    DefaultKeyMap,
		textInput: textinput.New(),
		paginator: paginator.New(),
		help:      help.New(),
		styleWidget: styleBase.Copy().
			Border(lipgloss.ThickBorder(), false, true, true, true).
			AlignHorizontal(lipgloss.Left).
			Height(10).
			Bold(true),
		styleStatePanel: styleBase.Copy().
			Align(lipgloss.Center),
	}

	am.textInput.Placeholder = "Enter File or URL to Search"
	am.textInput.SetValue(path)
	am.textInput.CharLimit = 1024
	am.textInput.Prompt = "┃"
	am.textInput.Focus()

	am.paginator.KeyMap = am.KeyMap.Paginator
	am.paginator.Type = paginator.Dots
	am.paginator.ActiveDot = styleBase.Copy().
		Bold(true).
		Render("•")
	am.paginator.InactiveDot = styleBase.Copy().
		Foreground(lipgloss.ANSIColor(termenv.ANSIBlack)).
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
			path := m.textInput.Value()
			if path == "" {
				return m, nil
			}
			m.searching = true
			return m, func() tea.Msg {
				switch {
				case strings.HasPrefix(path, "http"):
					m.fetcher.FetchByURL(m.response, APIBaseURL+APIByURLParameter+path)
				default:
					m.fetcher.FetchByFile(m.response, APIBaseURL, path)
				}

				return searchFinishedMsg{}
			}
		}

		if m.textInput.Focused() {
			switch {
			case key.Matches(msg, m.KeyMap.Focus):
				m.textInput.Blur()
				return m, nil
			}
			if !key.Matches(msg, m.KeyMap.Search) {
				m.textInput, cmd = m.textInput.Update(msg)
				return m, cmd
			}
		}

		switch {
		case key.Matches(msg, m.KeyMap.Help):
			m.help.ShowAll = !m.help.ShowAll
		case key.Matches(msg, m.KeyMap.AltScreen):
			m.altScreen = !m.altScreen
			if m.altScreen {
				return m, tea.EnterAltScreen
			}
			return m, tea.ExitAltScreen
		case key.Matches(msg, m.KeyMap.Focus):
			m.textInput.Focus()
		case key.Matches(msg, m.KeyMap.Paginator.NextPage, m.KeyMap.Paginator.PrevPage):
			m.paginator, cmd = m.paginator.Update(msg)
			return m, cmd
		}
	case tea.WindowSizeMsg:
		m.styleWidget.Width(msg.Width - 2)
		m.styleStatePanel.Width(msg.Width - 2)
		m.help.Width = msg.Width
		m.textInput.Width = msg.Width - 2
		return m, nil
	case searchFinishedMsg:
		m.paginator.SetTotalPages(len(m.response.Result))
		m.searching = false
	}

	return m, nil
}

func (m Model) View() string {
	s := m.textInput.View() + "\n"
	switch {
	case m.searching:
		s += m.styleWidget.Render(m.styleStatePanel.Render("⧖ SEARCHING ..."))
	case len(m.response.Result) == 0:
		s += m.styleWidget.Render(m.styleStatePanel.Render("✘ NO RESULTS ✘"))
	default:
		s += m.styleWidget.Render(m.styleStatePanel.Render(m.paginator.View()) + "\n" +
			styleCard.Render(m.response.Result[m.paginator.Page].View()) + "\n")
	}

	return s + "\n" +
		m.help.View(m.KeyMap)
}
