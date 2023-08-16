package config

import "github.com/charmbracelet/lipgloss"

type BorderConfig struct {
	Foreground lipgloss.Color
}

type AppearanceConfig struct {
	Background lipgloss.Color
	Foreground lipgloss.Color
	Border     BorderConfig
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

