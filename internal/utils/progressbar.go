package utils

import (
	"fmt"
	"github.com/schollz/progressbar/v3"
	"os"
)

func NewProgressWithBytes(bytes int64, description string) *progressbar.ProgressBar {
	return progressbar.NewOptions64(bytes,
		progressbar.OptionSetWriter(os.Stderr),
		progressbar.OptionOnCompletion(func() {
			_, _ = fmt.Fprint(os.Stderr, "\n")
		}),
		progressbar.OptionEnableColorCodes(true),
		progressbar.OptionShowBytes(true),
		progressbar.OptionShowCount(),
		progressbar.OptionSetWidth(100),
		progressbar.OptionSetDescription(description),
		progressbar.OptionSetTheme(progressbar.Theme{
			Saucer:        "[green]=[reset]",
			SaucerHead:    "[green]>[reset]",
			SaucerPadding: " ",
			BarStart:      "[",
			BarEnd:        "]",
		}))
}
