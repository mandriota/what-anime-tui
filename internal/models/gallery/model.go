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

const (
	APIBaseURL        = `https://api.trace.moe/search?anilistInfo`
	APIByURLParameter = `&url=`
)

var (
	styleBase = lipgloss.NewStyle().
			Background(config.Global.Appearance.Background).
			Foreground(config.Global.Appearance.Foreground)
	styleCard = lipgloss.NewStyle().
			PaddingLeft(4)
	styleGallery = lipgloss.NewStyle().
			Padding(1, 0, 1, 0)
)

type searchFinishedMsg struct {
	err error
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

	err error

	styleDeck       lipgloss.Style
	styleStatePanel lipgloss.Style
}

func New(path string) Model {
	am := Model{
		fetcher:  fetcher.New(),
		response: new(fetcher.Response),
		KeyMap:   DefaultKeyMap,
		spinner: spinner.New(spinner.WithSpinner(spinner.Spinner{
			Frames: ascii.ArtTelescope,
			FPS:    time.Millisecond * 1500 / time.Duration(len(ascii.ArtTelescope)),
		})),
		help: help.New(),
		styleDeck: styleBase.Copy().
			Border(lipgloss.ThickBorder(), true, false, false, false).
			BorderTopForeground(config.Global.Appearance.Border.Foreground).
			Height(10).
			Bold(true),
		styleStatePanel: styleBase.Copy().
			Align(lipgloss.Center),
	}

	am.textInput = textinput.New()
	am.textInput.Placeholder = "Enter File or URL to Search"
	am.textInput.Prompt = ""
	am.textInput.SetValue(path)
	am.textInput.Focus()

	am.paginator = paginator.New()
	am.paginator.KeyMap = am.KeyMap.Paginator
	am.paginator.Type = paginator.Dots
	am.paginator.ActiveDot = styleBase.Copy().
		Bold(true).
		Render("•")
	am.paginator.InactiveDot = styleBase.Copy().
		Faint(true).
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

			m.searching = true
			return m, tea.Sequence(
				m.spinner.Tick,
				func() tea.Msg {
					switch {
					case strings.HasPrefix(path, "http"):
						m.err = m.fetcher.FetchByURL(m.response, APIBaseURL+APIByURLParameter+path)
					default:
						m.err = m.fetcher.FetchByFile(m.response, APIBaseURL, path)
					}

					return searchFinishedMsg{err: m.err}
				})
		case key.Matches(msg, m.KeyMap.Blur):
			m.textInput.Blur()
			return m, nil
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
			m.textInput, cmd = m.textInput.Update(msg)
			return m, cmd
		}

		switch {
		case key.Matches(msg, m.KeyMap.Focus):
			m.textInput.Focus()
		case key.Matches(msg, m.KeyMap.Paginator.NextPage, m.KeyMap.Paginator.PrevPage):
			m.paginator, cmd = m.paginator.Update(msg)
			return m, cmd
		}
	case tea.WindowSizeMsg:
		m.styleDeck.Width(msg.Width)
		m.styleStatePanel.Width(msg.Width)
		m.help.Width = msg.Width
		m.textInput.Width = msg.Width - 1
		return m, nil
	case searchFinishedMsg:
		m.err = msg.err
		m.paginator.SetTotalPages(len(m.response.Result))
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

func (m Model) View() string {
	deck := ""
	switch {
	case m.err != nil:
		deck = m.styleStatePanel.Render(m.err.Error())
	case m.searching:
		deck = m.styleStatePanel.Render("SEARCHING\n" + m.spinner.View())
	case len(m.response.Result) == 0:
		deck = m.styleStatePanel.Render("NO RESULTS\n" + ascii.ArtNoResults)
	default:
		deck = m.styleStatePanel.Render(m.paginator.View()) + "\n" +
			styleCard.Render(m.response.Result[m.paginator.Page].View()+"\n")
	}

	return styleGallery.Render(m.textInput.View()+"\n"+
		m.styleDeck.Render(deck),
	) + "\n" +
		m.help.View(m.KeyMap)
}
