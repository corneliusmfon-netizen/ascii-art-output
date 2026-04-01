package main

import (
	"fmt"
	"os"
	"strings"
)

// The instruction for error handling is to print "ERROR: <message>" and exit if the user provides invalid input.
const usage = `Usage: go run . [OPTION] [STRING] [BANNER]

EX: go run . --output=<fileName.txt> something standard`

func main() {

	// Arguments handling to take everything the user types after the "go run ." part and process it.
	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println(usage)
		return
	}

	// Variables to hold sorted inforations.
	var outputFile, colorName, substring, input, bannerStyle string
	bannerStyle = "standard" // default

	var posArgs []string // A list to hold the positional arguments.

	// A walk through the arguments to find flags and peel them off.
	for _, arg := range args {
		if strings.HasPrefix(arg, "--output=") {
			outputFile = strings.TrimPrefix(arg, "--output=")
		} else if strings.HasPrefix(arg, "--color=") {
			colorName = strings.TrimPrefix(arg, "--color=")
		} else {
			posArgs = append(posArgs, arg)
		}
	}

	// A check for the number of positional arguments to determine which is which.
	if len(posArgs) == 1 {
		input = posArgs[0]
	} else if len(posArgs) == 2 {
		if colorName != "" {
			substring = posArgs[0]
			input = posArgs[1]
		} else {
			input = posArgs[0]
			bannerStyle = posArgs[1]
		}
	} else if len(posArgs) == 3 {
		substring = posArgs[0]
		input = posArgs[1]
		bannerStyle = posArgs[2]
	} else {
		fmt.Println(usage)
		return
	}

	// The color mapping to ANSI codes for terminal coloring.
	colors := map[string]string{
		"black": "\033[30m", "red": "\033[31m", "green": "\033[32m", "yellow": "\033[33m",
		"blue": "\033[34m", "magenta": "\033[35m", "cyan": "\033[36m", "white": "\033[37m",
	}

	var colorCode string
	reset := "\033[0m"

	// A check to see if the user provided a color and if it's valid, if so we get the ANSI code for it.
	if colorName != "" {
		if code, ok := colors[strings.ToLower(colorName)]; ok {
			colorCode = code
		} else {
			fmt.Println("ERROR: invalid color")
			return
		}
	}
	// The banner file reading part, we read the file corresponding to the banner style and handle any errors that may occur.
	data, err := os.ReadFile(bannerStyle + ".txt")
	if err != nil {
		fmt.Printf("ERROR: could not read banner file: %v\n", err)
		return
	}

	// File cleaning to handle different newline formats and splitting it into lines for easier access later.
	content := strings.ReplaceAll(string(data), "\r\n", "\n")
	lines := strings.Split(content, "\n")

	var finalResult strings.Builder

	// Handle Newlines in the input (e.g. "hello\nworld")
	words := strings.Split(input, "\\n")
	for _, word := range words {
		if word == "" {
			finalResult.WriteString("\n")
			continue
		}

		// Find Substring Locations
		var substrLoc [][2]int
		if substring != "" {
			for i := 0; i <= len(word)-len(substring); i++ {
				if word[i:i+len(substring)] == substring {
					substrLoc = append(substrLoc, [2]int{i, i + len(substring)})
				}
			}
		}

		// The main loop to construct the banner line by line, character by character, applying color if needed.
		for i := 1; i <= 8; i++ {
			for pos, char := range word {
				if char < 32 || char > 126 {
					continue
				}

				index := (int(char-32) * 9) + i
				part := lines[index]

				shouldColor := false
				if substring == "" && colorCode != "" {
					shouldColor = true
				} else if substring != "" {
					for _, loc := range substrLoc {
						if pos >= loc[0] && pos < loc[1] {
							shouldColor = true
							break
						}
					}
				}

				if shouldColor {
					finalResult.WriteString(colorCode + part + reset)
				} else {
					finalResult.WriteString(part)
				}
			}
			finalResult.WriteString("\n")
		}
	}

	// Output handling, if the user provided an output file we write to it, otherwise we print to the terminal.
	if outputFile != "" {
		err := os.WriteFile(outputFile, []byte(finalResult.String()), 0644)
		if err != nil {
			fmt.Printf("ERROR: could not write to file: %v\n", err)
			return
		}
	} else {
		fmt.Print(finalResult.String())
	}
}
