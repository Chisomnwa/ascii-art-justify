package art_justify

import (
	"log"
	"os"
	"strings"
)

// AsciiArt converts a given text into ASCII art using a banner file
func AsciiArt(text, banner string) [][]string {

	// Read banner file
	content, err := os.ReadFile(banner)
	if err != nil {
		log.Fatal("Error: ", err)
	}

	// Split banner into lines - f
	// Becomes slice of strings
	arts := strings.Split(string(content), "\n")

	// Convert escaped "\n" into real newline
	text = strings.ReplaceAll(text, "\\n", "\n")

	// Split input into lines 
	// Becomes a slice of strings
	// This helps in vertical layout of the split input words
	lines := strings.Split(text, "\n")

	var result [][]string

	for _, line := range lines {

		// Handle empty line (only ONE blank line, not 8 rows)
		if line == "" {
			empty := make([]string, 8) // create 8 empty row that serves as vertical gaps between words
			result = append(result, empty)
			continue
		}

		// split line into words
		// Helps in horizontal layout and removes extra spacing
		// Because we will control the spaces manually
		words := strings.Fields(line)

		for _, word := range words {
			rows := make([]string, 8)

			for row := 0; row < 8; row++ {
				var strBuild strings.Builder // To avoid repeated memory allocation due to repeated string concatenation

				for _, ch := range word {

					if ch < 32 || ch > 126 {
						continue
					}
					index := (int(ch) - 32) * 9 // Gets position of the seperator line for each character
					strBuild.WriteString(arts[index+row+1]) // Skips the seperator line here and moves down the character
				}

				rows[row] = strBuild.String() // Save each word, after building the full row for it
			}
			result = append(result, rows) // Each word becomes a slice of string
		}
	}

	return result
}
