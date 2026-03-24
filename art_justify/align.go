package art_justify

import (
	"fmt"
	"strings"
)

// LeftAlign prints ASCII art exactly as it is with no extra spacing
func LeftAlign(arts [][]string) {
	for row := 0; row < 8; row++ {
		lines := ""
		for _, char := range arts {
			lines += (char[row]) + " "
		}
		fmt.Println(lines)
	}
}

// RightAlign shifts the ASCII arts to the right by adding spaces to the left
func RightAlign(arts [][]string, total_width int) {
	for row := 0; row < 8; row++ {
		char_width := 0
		for _, char := range arts {
			char_width += len(char[row])
		}

		padding := total_width - char_width
		if padding < 0 {
			padding = 0
		}
		lines := ""
		for _, char := range arts {
			lines += char[row] + " "
		}
		fmt.Println(strings.Repeat(" ", padding) + lines)
	}
}

// CenterAlign places the ASCII art in the center of the terminal width
func CenterAlign(arts [][]string, total_width int) {
	for row := 0; row < 8; row++ {
		char_width := 0
		for _, char := range arts {
			char_width += len(char[row])
		}
		padding := (total_width - char_width) / 2
		if padding < 0 {
			padding = 0
		}
		lines := ""
		for _, char := range arts {
			lines += char[row] + " "
		}
		fmt.Println(strings.Repeat(" ", padding) + lines)
	}
}

// Justify places the spaces evenly between the ASCII art in the terminal
func JustifyAlign(arts [][]string, total_width int) {
	numWords := len(arts)
	gaps := numWords - 1

	// Calculate character width
	char_width := 0
	for _, char := range arts {
		if len(char) > 0 {
			char_width += len(char[0])
		}
	}

	// If only one word
	if gaps <= 0 {
		for i := 0; i < 8; i++ {
			fmt.Println(arts[0][i])
		}
		return
	}

	// Calculate spacing
	totalSpace := total_width - char_width - gaps
	if totalSpace < 0 {
		totalSpace = 0
	}

	baseSpace := totalSpace / gaps
	extraSpace := totalSpace % gaps

	// Build Output
	for row := 0; row < 8; row++ {
		var lines strings.Builder

		for i, char := range arts {
			lines.WriteString(char[row])

			if i < gaps {
				space := baseSpace + 1
				if i < extraSpace {
					space++
				}
				lines.WriteString(strings.Repeat(" ", space))
			}
		}
		fmt.Println(lines.String())
	}
}
