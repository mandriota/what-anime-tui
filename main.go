package main

import (
	"log"
	"os"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/mandriota/what-anime-tui/internal/models/gallery"
)

func main() {
	path := strings.Join(os.Args[1:], " ")

	p := tea.NewProgram(gallery.New(path))
	if _, err := p.Run(); err != nil {
		log.Fatalf("error while running UI: %v", err)
		os.Exit(1)
	}
}
