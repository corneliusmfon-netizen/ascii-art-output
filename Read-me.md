ASCII Art Master Machine

Welcome to the ASCII Art Master Machine! This is a professional-grade command-line tool written in Go that transforms regular text into giant, colorful ASCII art "banners."
🚀 Features

    3 Artistic Styles: Choose between standard, shadow, and thinkertoy fonts.

    Technicolor Support: Paint your art in 8 different colors (Red, Green, Blue, etc.).

    Highlighter Mode: Choose to color only a specific part (substring) of your word.

    File Storage: Save your masterpieces directly to a .txt file using the --output flag.

    Multi-Line Magic: Supports \n to draw art across multiple rows.

🛠️ How It Works

The program is built like a factory with four main departments:

    The Detective: Sorts through your commands and flags.

    The Librarian: Finds the correct font book and color codes.

    The Painter: Draws the art line-by-line (each letter is 8 lines tall!).

    The Postman: Delivers the art to your screen or saves it to a file.

📖 Usage Examples
1. Basic Art
code Bash

go run main.go "Hello World"

2. Changing Styles
code Bash

go run main.go "Awesome" shadow

3. Adding Color
code Bash

go run main.go --color=red "Fire"

4. Highlighting Specific Letters
code Bash

go run main.go --color=green "ll" "Hello"

5. Saving to a File
code Bash

go run main.go --output=my_banner.txt "Save Me" standard

📋 Requirements

    Go (Golang) installed.

    Font files (standard.txt, shadow.txt, thinkertoy.txt) must be in the same folder.

🌟 Author

Built with ❤️ by a future Software Engineer - Cornel!