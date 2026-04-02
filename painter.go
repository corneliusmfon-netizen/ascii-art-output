package main

import "strings"

const (
	// Each character in a banner font occupies exactly 9 lines
	// (one blank separator line + 8 lines of art).
	linesPerChar = 9

	// ANSI reset code — turns off all color/styling.
	colorReset = "\033[0m"
)

// RenderArt is The Painter. It takes the user's input text, an optional
// substring to highlight, the resolved ANSI color code, and the pre-loaded
// font lines, then builds the complete ASCII art string and returns it.
func RenderArt(input, substring, colorCode string, lines []string) string {
	var result strings.Builder

	// Support \n in the input to render multi-line banners.
	words := strings.Split(input, "\\n")

	for _, word := range words {
		// An empty segment between \n markers means a blank line.
		if word == "" {
			result.WriteString("\n")
			continue
		}

		// Find every position where the substring occurs in this word
		// so we know which characters to colorize.
		substrRanges := findSubstringRanges(word, substring)

		// Each ASCII character is 8 art-lines tall. We iterate row by row
		// (rows 1–8), printing every character's slice for that row.
		for row := 1; row <= 8; row++ {
			for pos, char := range word {
				// Skip non-printable / out-of-range characters.
				if char < 32 || char > 126 {
					continue
				}

				// Formula: the font file stores characters starting from
				// ASCII 32 (space). Each block is 9 lines (1 blank + 8 art).
				lineIndex := (int(char-32) * linesPerChar) + row
				part := lines[lineIndex]

				if shouldColorize(pos, substring, colorCode, substrRanges) {
					result.WriteString(colorCode + part + colorReset)
				} else {
					result.WriteString(part)
				}
			}
			result.WriteString("\n")
		}
	}

	return result.String()
}

// findSubstringRanges returns a list of [start, end) index pairs for every
// occurrence of substr inside word. Returns nil when substr is empty.
func findSubstringRanges(word, substr string) [][2]int {
	if substr == "" {
		return nil
	}

	var ranges [][2]int
	for i := 0; i <= len(word)-len(substr); i++ {
		if word[i:i+len(substr)] == substr {
			ranges = append(ranges, [2]int{i, i + len(substr)})
		}
	}

	return ranges
}

// shouldColorize returns true when the character at pos should be painted
// with the chosen color, based on the coloring rules:
//   - No substring specified → colorize everything (if a color was given).
//   - Substring specified   → colorize only the characters inside a match.
func shouldColorize(pos int, substring, colorCode string, ranges [][2]int) bool {
	if colorCode == "" {
		return false
	}

	// Color the whole word when no specific substring was requested.
	if substring == "" {
		return true
	}

	// Color only the characters that fall inside a matched substring range.
	for _, r := range ranges {
		if pos >= r[0] && pos < r[1] {
			return true
		}
	}

	return false
}
