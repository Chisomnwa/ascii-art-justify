package art_justify

import (
	"os"
	"os/exec"
	"strconv"
	"strings"
)

// get terminal width of the current terminal
func GetTerminalWidth() int{
	cmd := exec.Command("stty", "size")
	cmd.Stdin = os.Stdin // attach command to current terminal

	output, err := cmd.Output()
	if err != nil {
		return 80 //fallback width(default)
	}

	// split "rows columns"
	parts := strings.Fields((string(output)))
	if len(parts) != 2 {
		return 80
	}

	//  converting width to integer
	total_width, err := strconv.Atoi(parts[1])
	if err != nil {
		return 80
	}

	return total_width
}