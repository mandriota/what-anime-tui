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

type appearance struct {
	Background lipgloss.Color
	Foreground lipgloss.Color
	Border     border
}

var Global = struct {
	Appearance appearance
}{
	Appearance: appearance{
		Background: lipgloss.Color(termenv.ANSICyan.String()),
		Foreground: lipgloss.Color(termenv.ANSIWhite.String()),
		Border: border{
			Foreground: lipgloss.Color(termenv.ANSIBlue.String()),
		},
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
