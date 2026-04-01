package art_justify

import (
	"fmt"
	"strings"
	 "os"
)

func LeftAlign(text_art string) {
	fmt.Print(text_art)

}

func RightAlign(text_art string, width int) {
	lines := strings.Split(strings.TrimRight(text_art, "\n"), "\n")

	for _, line := range lines {
		padding := (width - len(line))
		if padding < 0 {
			padding = 0
		}
		fmt.Printf("%s%s\n",strings.Repeat(" ", padding), line)
	}
}

func CenterAlign(text_art string, width int) {
	lines := strings.Split(strings.TrimRight(text_art, "\n"), "\n")

	for _, line := range lines {
		padding := (width - len(line))/2
		if padding < 0 {
			padding = 0
		}
		fmt.Printf("%s%s\n",strings.Repeat(" ", padding), line)
	}
}

func JustifyAlign(text_art string, text string, width int, banner string) {
	lines := strings.Split(strings.TrimRight(text_art, "\n"), "\n")
	words := strings.Fields(text)
	gaps := len(words) - 1

	if gaps <= 0 {
		fmt.Println(text_art)
		return
	}

	artWidth := 0
	for _, line := range lines {
		if len(line) > artWidth {
			artWidth = len(line)
		}
	}

	padding := width - artWidth
	baseSpace := padding / gaps
	remainder := padding % gaps

	spaceWidth := getSpaceWidth(banner)
	spaceSep := strings.Repeat(" ", spaceWidth) // the actual gap between words

	for _, line := range lines {
		parts := strings.Split(line, spaceSep)
		result := ""
		gapIndex := 0
		for i, part := range parts {
			result += part
			if i < len(parts)-1 {
				extra := 0
				if gapIndex < remainder {
					extra = 1
				}
				result += strings.Repeat(" ", spaceWidth+baseSpace+extra)
				gapIndex++
			}
		}
		fmt.Println(result)
	}
}

func getSpaceWidth(banner string) int {
	content, err := os.ReadFile(banner)
	if err != nil {
		return 8 // default fallback
	}
	arts := strings.Split(string(content), "\n")
	return len(arts[1]) // space is ASCII 32, index 0 in the banner, row 1 = arts[1]
}