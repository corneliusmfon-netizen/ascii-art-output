package main

import (
	"fmt"
	"os"
)

const usage = `Usage: go run . [OPTION] [STRING] [BANNER]
EX: go run . --output=<fileName.txt> something standard`

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println(usage)
		return
	}

	// 1. THE DETECTIVE — parse and validate all CLI arguments.
	cfg, err := ParseArgs(args)
	if err != nil {
		fmt.Printf("ERROR: %v\n", err)
		fmt.Println(usage)
		return
	}

	// 2. THE LIBRARIAN — load the font file and resolve the color code.
	lines, err := LoadBannerLines(cfg.BannerStyle)
	if err != nil {
		fmt.Printf("ERROR: %v\n", err)
		return
	}

	colorCode, err := ResolveColor(cfg.ColorName)
	if err != nil {
		fmt.Printf("ERROR: %v\n", err)
		return
	}

	// 3. THE PAINTER — render the ASCII art.
	result := RenderArt(cfg.Input, cfg.Substring, colorCode, lines)

	// 4. THE POSTMAN — deliver to screen or file.
	if err := Deliver(result, cfg.OutputFile); err != nil {
		fmt.Printf("ERROR: %v\n", err)
	}
}
