package ascii

import (
	_ "embed"
	"strings"
)

var (
	//go:embed assets/searching.txt
	artTelescope string

	ArtTelescope = strings.Split(artTelescope, "===\n")

	//go:embed assets/no_results.txt
	ArtNoResults string
)
