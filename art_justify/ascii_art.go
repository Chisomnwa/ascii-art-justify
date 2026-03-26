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

	// Split banner into lines
	arts := strings.Split(string(content), "\n")

	// Convert escaped "\n" into real newline
	text = strings.ReplaceAll(text, "\\n", "\n")

	// Split input into lines
	lines := strings.Split(text, "\n")

	var result [][]string

	for _, line := range lines {

		// Handle empty line (only ONE blank line, not 8 rows)
		if line == "" {
			empty := make([]string, 8) // create 8 empty row
			result = append(result, empty)
			continue
		}

		// split line into words
		words := strings.Fields(line)

		for _, word := range words {
			rows := make([]string, 8)

			for row := 0; row < 8; row++ {
				var strBuild strings.Builder

				for _, ch := range word {

					if ch < 32 || ch > 126 {
						continue
					}
					index := (int(ch) - 32) * 9
					strBuild.WriteString(arts[index+row+1])
				}

				rows[row] = strBuild.String()
			}
			result = append(result, rows)
		}
	}

	return result
}
