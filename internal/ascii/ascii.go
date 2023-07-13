package ascii

import _ "embed"

var (
	//go:embed assets/searching.txt
	ArtTelescope string

	//go:embed assets/no_results.txt
	ArtNoResults string
)
