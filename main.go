package main

import (
	"log"
	"os"
	"path/filepath"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/mandriota/what-anime-tui/internal/config"
	"github.com/mandriota/what-anime-tui/internal/models/gallery"
	"github.com/muesli/termenv"
	"github.com/pelletier/go-toml/v2"
)

var cfg = config.GeneralConfig{
	Appearance: config.AppearanceConfig{
		Background: lipgloss.Color(termenv.ANSICyan.String()),
		Foreground: lipgloss.Color(termenv.ANSIWhite.String()),
		Border: config.BorderConfig{
			Foreground: lipgloss.Color(termenv.ANSIWhite.String()),
		},
	},
	Fetcher: config.FetcherConfig{
		ApiUrlByUrl: `https://api.trace.moe/search?anilistInfo&url={{ .Path }}`,
		ApiUrlByFile: `https://api.trace.moe/search?anilistInfo`,
	},
}

func init() {
	cfgDir, err := os.UserHomeDir()
	if err != nil {
		return
	}

	path := filepath.Join(cfgDir, ".config", "wat", "wat.toml")
	fs, err := os.Open(path)
	if os.IsNotExist(err) {
		return
	}
	if err != nil {
		log.Println("failed to read configuration file")
	}
	defer fs.Close()

	if err := toml.NewDecoder(fs).Decode(&cfg); err != nil {
		log.Println(err)
	}
}

func main() {
	path := strings.Join(os.Args[1:], " ")

	p := tea.NewProgram(gallery.New(cfg, path))
	if _, err := p.Run(); err != nil {
		log.Fatalf("error while running UI: %v", err)
		os.Exit(1)
	}
}
