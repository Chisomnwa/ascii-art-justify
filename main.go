package main

import (
	"ascii-justify/art_justify"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 || len(os.Args) > 4 {
		log.Fatal("\nUsage: go run . [OPTION] [STRING] [BANNER]\n\nExample: go run . --align=right something standard")
	}

	text := ""
	banner := "standard.txt"
	alignment := "left"

	if len(os.Args) == 2 {
		text = os.Args[1]

		if text == "" {
			return
		}

		if text == "\\n" {
			fmt.Println()
			return
		}

	} else if len(os.Args) == 3 {
		arg := os.Args[1]

		if strings.HasPrefix(arg, "--align=") {
			alignment = strings.TrimPrefix(arg, "--align=")
			text = os.Args[2]
		} else {
			text = os.Args[1]
			banner = os.Args[2] + ".txt"
		}

	} else if len(os.Args) == 4 {
		arg := os.Args[1]

		if !strings.HasPrefix(arg, "--align=") {
			log.Fatal("Invalid align flag. Expected format: --align=left|right|center|justify")
		}

		alignment = strings.ToLower(strings.TrimPrefix(arg, "--align="))
		text = os.Args[2]
		banner = os.Args[3] + ".txt"
	}

	// get terminal width for alignment calculations
	total_width := art_justify.GetTerminalWidth()
	if len(text) > total_width {
		log.Fatal("Terminal size exceeded")
	}

	// Generate ASCII
	result := art_justify.AsciiArt(text, banner)

	// Applying the alignment
	switch alignment {
	case "left":
		art_justify.LeftAlign(result)
	case "right":
		art_justify.RightAlign(result, total_width)
	case "center":
		art_justify.CenterAlign(result, total_width)
	case "justify" :
		art_justify.JustifyAlign(result, total_width)

	default:
		(log.Fatal("error"))
	}

}
