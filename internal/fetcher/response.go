package fetcher

import (
	"fmt"
	"strings"
)

type ResponseResult struct {
	Anilist struct {
		// Id    int
		// IdMal int
		Title struct {
			Native  string
			Romaji  string
			English string
		}
		// Synonyms []string
		IsAdult bool
	}
	// Filename   string
	Episode    any
	From       float64
	To         float64
	Similarity float64
	// Video      string
	// Image 	  string

	sb *strings.Builder
}

func (rr ResponseResult) View() string {
	if rr.sb != nil {
		return rr.sb.String()
	}
	rr.sb = &strings.Builder{}

	for _, title := range [...][2]string{
		{"    Native:\t%s\n", rr.Anilist.Title.Native},
		{"    Romaji:\t%s\n", rr.Anilist.Title.Romaji},
		{"   English:\t%s\n", rr.Anilist.Title.English},
	} {
		if title[1] == "" {
			title[1] = "???"
		}
		fmt.Fprintf(rr.sb, title[0], title[1])
	}
	if rr.Anilist.IsAdult {
		fmt.Fprintf(rr.sb, "  Is Adult:\tYes\n")
	} else {
		fmt.Fprintf(rr.sb, "  Is Adult:\tNo\n")
	}
	if rr.Episode == nil {
		rr.Episode = "???"
	}
	fmt.Fprintf(rr.sb, "   Episode:\t%v\n", rr.Episode)
	fmt.Fprintf(rr.sb, "      From:\t%s\n", formatSecondsToDuration(rr.From))
	fmt.Fprintf(rr.sb, "        To:\t%s\n", formatSecondsToDuration(rr.To))

	progressBar := "Very Low"
	switch {
	case rr.Similarity >= 0.95:
		progressBar = "Very High"
	case rr.Similarity >= 0.9:
		progressBar = "High"
	case rr.Similarity >= 0.8:
		progressBar = "Low"
	}

	fmt.Fprintf(rr.sb, "Similarity:\t%.2f%% (%s)", rr.Similarity*100, progressBar)

	return rr.sb.String()
}

type Response struct {
	// FrameCount int
	Error  string
	Result []ResponseResult
}
