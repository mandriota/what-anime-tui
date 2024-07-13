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
package main

import (
	"flag"
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
		ApiUrlByUrl:  `https://api.trace.moe/search?anilistInfo&url={{ .Path }}`,
		ApiUrlByFile: `https://api.trace.moe/search?anilistInfo`,
	},
}

var cliQuit = false

func init() {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		homeDir = "/"
	}

	cfgDir := filepath.Join(homeDir, ".config", "wat", "wat.toml")

	flag.StringVar(&cfgDir, "c", cfgDir, "path to the configuration file")
	flag.BoolVar(&cliQuit, "q", false, "if set, program will be closed after finding first result")
	flag.Parse()

	fs, err := os.Open(cfgDir)
	if os.IsNotExist(err) {
		return
	}
	if err != nil {
		log.Println("failed to read the configuration file")
		return
	}
	defer fs.Close()

	if err := toml.NewDecoder(fs).Decode(&cfg); err != nil {
		log.Println(err)
	}
}

func main() {
	path := strings.Join(flag.Args(), " ")

	p := tea.NewProgram(gallery.New(cfg, cliQuit, path))
	if _, err := p.Run(); err != nil {
		log.Fatalln("error while running UI: ", err)
	}
}
