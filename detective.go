package main

import (
	"errors"
	"strings"
)

// Config holds all the information parsed from the user's CLI arguments.
type Config struct {
	OutputFile  string // destination file, e.g. "banner.txt"
	ColorName   string // e.g. "red", "green"
	Substring   string // the part of the input to colorize
	Input       string // the text to render as ASCII art
	BannerStyle string // "standard", "shadow", or "thinkertoy"
}

// ParseArgs is The Detective. It walks through every argument the user
// provided, peels off the flags (--output, --color), then assigns the
// remaining positional arguments to Input, Substring, and BannerStyle.
func ParseArgs(args []string) (*Config, error) {
	cfg := &Config{
		BannerStyle: "standard", // sensible default
	}

	var posArgs []string

	for _, arg := range args {
		switch {
		case strings.HasPrefix(arg, "--output="):
			cfg.OutputFile = strings.TrimPrefix(arg, "--output=")

		case strings.HasPrefix(arg, "--color="):
			cfg.ColorName = strings.TrimPrefix(arg, "--color=")

		default:
			posArgs = append(posArgs, arg)
		}
	}

	// Decide which positional slot is which based on how many there are
	// and whether a --color flag was also supplied.
	switch len(posArgs) {
	case 1:
		// go run . "Hello"
		cfg.Input = posArgs[0]

	case 2:
		if cfg.ColorName != "" {
			// go run . --color=red "ll" "Hello"  ← substring + input
			cfg.Substring = posArgs[0]
			cfg.Input = posArgs[1]
		} else {
			// go run . "Hello" shadow  ← input + banner style
			cfg.Input = posArgs[0]
			cfg.BannerStyle = posArgs[1]
		}

	case 3:
		// go run . "ll" "Hello" shadow  ← substring + input + banner style
		cfg.Substring = posArgs[0]
		cfg.Input = posArgs[1]
		cfg.BannerStyle = posArgs[2]

	default:
		return nil, errors.New("unexpected number of arguments")
	}

	return cfg, nil
}
