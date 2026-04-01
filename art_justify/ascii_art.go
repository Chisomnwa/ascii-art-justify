package art_justify

import (
	"log"
	"os"
	"strings"
)

/*
AsciiArt converts a given text into ASCII art using a banner file

Args:

	text: tis is the CLI input string
	banner: file that contains the ascii art characters

Returns:

	a slice of strings (ascii art representation of the input string)
*/
func AsciiArt(text, banner string) string {
	// Read the banner file
	content, err := os.ReadFile(banner)
	if err != nil {
		log.Print(err.Error())
		return ""
	}

	arts := strings.Split(string(content), "\n")

	text = strings.ReplaceAll(text, "\\n", "\n")

	words := strings.Split(text, "\n")

	var result strings.Builder

	for _, word := range words {
		if word == "" {
			result.WriteString("\n")
			continue
		}

		for row := 0; row < 8; row++ {
			for _, char:= range word {
				index := (int(char)-32) * 9
				result.WriteString(arts[index+row+1])
			}
			result.WriteString("\n")
		}
	}
	return result.String()

}
