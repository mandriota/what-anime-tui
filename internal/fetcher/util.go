package fetcher

import (
	"fmt"
	"math"
)

func formatSecondsToDuration(s float64) string {
	return fmt.Sprintf("%02d:%02d:%02.1f",
		int(s/3600),
		int(s/60)%60,
		math.Mod(s, 60),
	)
}
