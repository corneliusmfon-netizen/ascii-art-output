package main

import (
	"fmt"
	"os"
)

// Deliver is The Postman. It receives the finished ASCII art string and
// decides where to send it:
//   - If outputFile is empty, it prints to the terminal (stdout).
//   - If outputFile is provided, it writes the content to that file,
//     creating it if it doesn't exist and overwriting it if it does.
func Deliver(content, outputFile string) error {
	if outputFile == "" {
		fmt.Print(content)
		return nil
	}

	err := os.WriteFile(outputFile, []byte(content), 0644)
	if err != nil {
		return fmt.Errorf("could not write to file '%s': %w", outputFile, err)
	}

	fmt.Printf("Art saved to '%s' successfully.\n", outputFile)
	return nil
}
