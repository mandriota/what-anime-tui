package config

import (
	"log"
	"os"
	"path/filepath"

	"github.com/charmbracelet/lipgloss"
	"github.com/muesli/termenv"
	"github.com/pelletier/go-toml/v2"
)

type border struct {
	Foreground lipgloss.Color
}

type AppearanceConfig struct {
	Background lipgloss.Color
	Foreground lipgloss.Color
	Border     border
}

// Warning: this fetcher config is not stable and may be changed in future releases.
type FetcherConfig struct {
	ApiUrlByUrl  string
	ApiUrlByFile string
}

type GeneralConfig struct {
	Appearance AppearanceConfig
	Fetcher FetcherConfig
}

var Global = GeneralConfig{
	Appearance: AppearanceConfig{
		Background: lipgloss.Color(termenv.ANSICyan.String()),
		Foreground: lipgloss.Color(termenv.ANSIWhite.String()),
		Border: border{
			Foreground: lipgloss.Color(termenv.ANSIWhite.String()),
		},
	},
	Fetcher: FetcherConfig{
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

	if err := toml.NewDecoder(fs).Decode(&Global); err != nil {
		log.Println(err)
	}
}
