package main

import (
	"fmt"
	"os"
	"strings"
)

// colorMap is the Librarian's internal catalog of supported colors.
// Each name maps to its ANSI escape code for terminal coloring.
var colorMap = map[string]string{
	"black":   "\033[30m",
	"red":     "\033[31m",
	"green":   "\033[32m",
	"yellow":  "\033[33m",
	"blue":    "\033[34m",
	"magenta": "\033[35m",
	"cyan":    "\033[36m",
	"white":   "\033[37m",
}

// LoadBannerLines is part of The Librarian's job. It opens the correct
// font ".txt" file for the requested banner style, normalizes line endings,
// and returns every line as a slice so the Painter can index into it later.
func LoadBannerLines(bannerStyle string) ([]string, error) {
	data, err := os.ReadFile(bannerStyle + ".txt")
	if err != nil {
		return nil, fmt.Errorf("could not read banner file '%s.txt': %w", bannerStyle, err)
	}

	// Normalize Windows-style CRLF to Unix LF before splitting.
	content := strings.ReplaceAll(string(data), "\r\n", "\n")
	lines := strings.Split(content, "\n")

	return lines, nil
}

// ResolveColor is the other half of The Librarian's job. It looks up
// the ANSI code for the requested color name. An empty name is valid
// (it means "no coloring"), and returns an empty string without an error.
func ResolveColor(colorName string) (string, error) {
	if colorName == "" {
		return "", nil
	}

	code, ok := colorMap[strings.ToLower(colorName)]
	if !ok {
		return "", fmt.Errorf("invalid color '%s': supported colors are black, red, green, yellow, blue, magenta, cyan, white", colorName)
	}

	return code, nil
}
