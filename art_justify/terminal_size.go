package art_justify

import (
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

// get terminal width of the current terminal
func GetTerminalWidth() int {
	cmd := exec.Command("stty", "size")
	cmd.Stdin = os.Stdin

	output, err := cmd.Output()
	if err != nil {
		log.Print(err.Error())
		return 80 // Default size
	}

	height_Width := strings.Fields(string(output))

	width, err := strconv.Atoi(height_Width[1])
	if err != nil {
		log.Print(err.Error())
		return 80
	}

	return width
}
