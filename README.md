# ASCII Art Master Machine

A command-line tool written in Go that transforms plain text into giant, styled ASCII art banners — with color support, substring highlighting, multi-line rendering, and file output.

```
 _    _      _ _
| |  | |    | | |
| |__| | ___| | | ___
|  __  |/ _ \ | |/ _ \
| |  | |  __/ | | (_) |
|_|  |_|\___|_|_|\___/
```

---

## Table of Contents

- [Overview](#overview)
- [Project Structure](#project-structure)
- [Features](#features)
- [Requirements](#requirements)
- [Installation](#installation)
- [Usage](#usage)
  - [Basic Usage](#basic-usage)
  - [Changing Font Styles](#changing-font-styles)
  - [Adding Color](#adding-color)
  - [Highlighting a Substring](#highlighting-a-substring)
  - [Multi-line Input](#multi-line-input)
  - [Saving to a File](#saving-to-a-file)
  - [Combining Flags](#combining-flags)
- [Available Font Styles](#available-font-styles)
- [Available Colors](#available-colors)
- [How It Works](#how-it-works)
  - [The Detective — detective.go](#the-detective--detectivego)
  - [The Librarian — librarian.go](#the-librarian--librariango)
  - [The Painter — painter.go](#the-painter--paintergo)
  - [The Postman — postman.go](#the-postman--postmango)
  - [The Manager — main.go](#the-manager--maingo)
- [Error Handling](#error-handling)
- [Example Output](#example-output)
- [Author](#author)

---

## Overview

ASCII Art Master Machine reads text you provide on the command line and renders it as large ASCII art using pre-built font files. Each character is drawn as an 8-line-tall block of symbols, assembled side by side to spell out your message.

The program supports three distinct artistic styles, eight terminal colors, optional substring highlighting (where only part of the word is colored), multi-line output via `\n`, and direct file export.

---

## Project Structure

The codebase is split into five files, each with a single, well-defined responsibility:

```
ascii-art/
├── main.go          # Wires the four departments together — the manager
├── detective.go     # Parses and validates all CLI arguments
├── librarian.go     # Loads the font file and resolves color codes
├── painter.go       # Renders the ASCII art character by character
├── postman.go       # Delivers output to the terminal or a file
├── standard.txt     # Font file — classic block letters
├── shadow.txt       # Font file — letters with a shadow effect
└── thinkertoy.txt   # Font file — playful, open letter style
```

All five `.go` files belong to `package main`, which means Go treats them as a single program during compilation. No internal imports are needed — the functions defined in `detective.go`, `librarian.go`, `painter.go`, and `postman.go` are all directly accessible from `main.go`.

---

## Features

- **3 font styles** — `standard`, `shadow`, and `thinkertoy`
- **8 terminal colors** — black, red, green, yellow, blue, magenta, cyan, white
- **Substring highlighting** — color only a specific part of the word, leave the rest plain
- **Multi-line support** — use `\n` in your input to draw art across multiple rows
- **File output** — save the rendered art to a `.txt` file with `--output=`
- **Descriptive errors** — clear `ERROR:` messages when input is invalid

---

## Requirements

- [Go](https://go.dev/dl/) 1.18 or later installed on your machine
- The three font files (`standard.txt`, `shadow.txt`, `thinkertoy.txt`) must be present in the same directory as the `.go` files

---

## Installation

1. Clone or download the repository into a local folder:

```bash
git clone https://github.com/your-username/ascii-art.git
cd ascii-art
```

2. Confirm Go is installed:

```bash
go version
```

3. No additional dependencies are required. All packages used (`fmt`, `os`, `strings`) are part of Go's standard library.

---

## Usage

The general syntax is:

```bash
go run . [--output=<file>] [--color=<color>] [substring] <text> [banner]
```

Square brackets `[ ]` denote optional parts. Arguments without brackets are required.

| Part | Description |
|---|---|
| `--output=<file>` | Optional. Save output to a file instead of printing it |
| `--color=<color>` | Optional. Color the entire text, or only the substring if one is given |
| `substring` | Optional. The part of the text to colorize (requires `--color`) |
| `text` | Required. The string to render as ASCII art |
| `banner` | Optional. Font style: `standard`, `shadow`, or `thinkertoy`. Defaults to `standard` |

---

### Basic Usage

Render a word using the default `standard` font:

```bash
go run . "Hello"
```

---

### Changing Font Styles

Pass a banner style as the second positional argument:

```bash
go run . "Hello" shadow
go run . "Hello" thinkertoy
go run . "Hello" standard
```

---

### Adding Color

Use the `--color=` flag to color the entire output:

```bash
go run . --color=red "Fire"
go run . --color=cyan "Ice" shadow
go run . --color=yellow "Sunshine" thinkertoy
```

---

### Highlighting a Substring

When `--color` is combined with a substring argument, only the matching characters are colored. The rest of the word is printed in plain terminal text.

The substring must be the first positional argument, and the full text must be the second:

```bash
go run . --color=green "ll" "Hello"
```

In this example, only the two `l` characters in "Hello" will appear in green. "He" and "o" remain uncolored.

If the substring appears more than once in the text, every occurrence is colored.

---

### Multi-line Input

Use the literal characters `\n` in your input string to produce multi-line ASCII art:

```bash
go run . "Hello\nWorld"
go run . --color=blue "Go\nLang" shadow
```

An empty segment between `\n` markers (e.g. `"Hello\n\nWorld"`) produces a blank line in the output.

---

### Saving to a File

Use the `--output=` flag with a filename to write the result to disk instead of printing it:

```bash
go run . --output=banner.txt "Hello" standard
go run . --output=result.txt --color=red "Error"
```

When saving to a file, the ANSI color codes are written as-is. Note that some text editors may not render them visually — they are designed for terminal display.

A confirmation message is printed to the terminal when the file is saved successfully.

---

### Combining Flags

Flags can be combined freely. Order of flags does not matter — they are always peeled off first before positional arguments are interpreted:

```bash
go run . --color=magenta --output=out.txt "ll" "Hello World" shadow
```

This renders "Hello World" in the `shadow` style, colors every occurrence of "ll" in magenta, and saves the result to `out.txt`.

---

## Available Font Styles

| Style | Description |
|---|---|
| `standard` | Classic solid block letters. The default if no style is specified |
| `shadow` | Block letters with a shadow effect on the lower-right |
| `thinkertoy` | A lighter, more open style built from simpler characters |

The font files must be in the same directory as the program. If a file is missing, the program exits with a descriptive error message.

---

## Available Colors

| Name | Terminal effect |
|---|---|
| `black` | Dark / invisible on dark terminals |
| `red` | Bright red |
| `green` | Bright green |
| `yellow` | Bright yellow |
| `blue` | Bright blue |
| `magenta` | Bright magenta / pink |
| `cyan` | Bright cyan / teal |
| `white` | Bright white |

Color names are case-insensitive — `Red`, `RED`, and `red` are all accepted. Any color name not in this list will cause the program to exit with `ERROR: invalid color`.

---

## How It Works

The program is divided into four workers coordinated by a single manager. Here is what each file does:

### The Detective — `detective.go`

`ParseArgs(args []string) (*Config, error)`

Reads every word the user typed after `go run .` and sorts it into a structured `Config` object with labeled fields: `OutputFile`, `ColorName`, `Substring`, `Input`, and `BannerStyle`.

Flags (anything starting with `--`) are peeled off first. The remaining positional arguments are then interpreted based on how many there are and whether `--color` was also provided:

- 1 positional argument → it is the text to render
- 2 positional arguments + `--color` → first is the substring, second is the text
- 2 positional arguments, no `--color` → first is the text, second is the banner style
- 3 positional arguments → substring, text, banner style

Any other count returns an error.

---

### The Librarian — `librarian.go`

`LoadBannerLines(bannerStyle string) ([]string, error)`

Opens the font file corresponding to the chosen banner style (e.g. `standard.txt`), normalizes Windows-style line endings (`\r\n` → `\n`), and returns every line of the file as a slice of strings so the Painter can index into it later.

`ResolveColor(colorName string) (string, error)`

Looks up the color name in a map and returns the corresponding ANSI escape code (e.g. `"red"` → `"\033[31m"`). An empty color name returns an empty string without error — meaning no coloring will be applied.

---

### The Painter — `painter.go`

`RenderArt(input, substring, colorCode string, lines []string) string`

The most complex component. It takes the input text, splits it on `\n` markers, and for each segment renders the ASCII art row by row.

Each character in the font file occupies exactly 9 lines: one blank separator line followed by 8 lines of art. The line index for a given character and row is calculated as:

```
lineIndex = (asciiCode - 32) × 9 + rowNumber
```

The Painter iterates through rows 1 to 8, and for each row it writes every character's art slice side by side — producing one line of terminal output per iteration. This repeats eight times to complete each letter block.

Two helper functions support it:

- `findSubstringRanges` — scans the word and records the start and end positions of every occurrence of the substring
- `shouldColorize` — for each character position, decides whether to wrap it in the color code and reset sequence

---

### The Postman — `postman.go`

`Deliver(content, outputFile string) error`

The simplest component. If `outputFile` is an empty string, it prints the rendered art to standard output with `fmt.Print`. If a filename was given, it writes the content to disk using `os.WriteFile` with `0644` permissions (owner can read and write; others can only read), then prints a confirmation message.

---

### The Manager — `main.go`

The entry point. It reads `os.Args`, checks that at least one argument was provided, then calls the four workers in order:

1. `ParseArgs` → get the config
2. `LoadBannerLines` → get the font data
3. `ResolveColor` → get the ANSI color code
4. `RenderArt` → produce the art string
5. `Deliver` → send it to the right destination

Every step that returns an error is checked immediately. If any worker fails, the program prints `ERROR: <message>` and exits — no partial output is produced.

---

## Error Handling

The program follows a consistent error pattern: print `ERROR: <descriptive message>` to stdout and exit without producing any art output.

| Situation | Error message |
|---|---|
| No arguments provided | Prints usage instructions |
| Wrong number of positional arguments | Prints usage instructions |
| Unrecognized color name | `ERROR: invalid color '<name>'` |
| Font file not found | `ERROR: could not read banner file '<name>.txt'` |
| Cannot write to output file | `ERROR: could not write to file '<name>'` |

---

## Example Output

Running:

```bash
go run . "Hi" shadow
```

Produces something like:

```
 _  _   _
| || | (_)
| __ |  _
|_||_| |_|
```

Running:

```bash
go run . --color=cyan "Go" thinkertoy
```

Produces the word "Go" in thinkertoy style, rendered in cyan in the terminal.

---

## Author

Built by **Cornel and Oluwatosin**

> "The best code is code that each piece clearly knows its own job."
