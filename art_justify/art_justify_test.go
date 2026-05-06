package art_justify

import (
	"bytes"
	"os"
	"strings"
	"testing"
)

// helper: capture stdout from functions that print
func captureOutput(f func()) string {
	// backup real stdout
	old := os.Stdout

	r, w, _ := os.Pipe()
	os.Stdout = w

	f()

	w.Close()
	os.Stdout = old

	var buf bytes.Buffer
	buf.ReadFrom(r)

	return buf.String()
}

// -------------------------
// ASCII ART TESTS
// -------------------------

func TestAsciiArt_NotEmpty(t *testing.T) {
	result := AsciiArt("A", "../standard.txt")

	if len(result) == 0 {
		t.Errorf("Expected non-empty result")
	}

	if len(result[0]) != 8 {
		t.Errorf("Expected 8 rows per character, got %d", len(result[0]))
	}
}

func TestAsciiArt_NewLineHandling(t *testing.T) {
	result := AsciiArt("A\\nB", "../standard.txt")

	if len(result) < 2 {
		t.Errorf("Expected multiple blocks for newline input")
	}
}

// -------------------------
// ALIGNMENT TESTS
// -------------------------

func TestLeftAlign(t *testing.T) {
	arts := [][]string{
		{"A", "A", "A", "A", "A", "A", "A", "A"},
		{"B", "B", "B", "B", "B", "B", "B", "B"},
	}

	output := captureOutput(func() {
		LeftAlign(arts)
	})

	if !strings.Contains(output, "A B") {
		t.Errorf("Expected left aligned output to contain 'A B'")
	}
}

func TestRightAlign(t *testing.T) {
	arts := [][]string{
		{"A", "A", "A", "A", "A", "A", "A", "A"},
		{"B", "B", "B", "B", "B", "B", "B", "B"},
	}

	width := 20

	output := captureOutput(func() {
		RightAlign(arts, width)
	})

	lines := strings.Split(strings.TrimSpace(output), "\n")

	for _, line := range lines {
		if len(line) > width {
			t.Errorf("Line exceeds terminal width")
		}
	}
}

func TestCenterAlign(t *testing.T) {
	arts := [][]string{
		{"AA", "AA", "AA", "AA", "AA", "AA", "AA", "AA"},
	}

	width := 20

	output := captureOutput(func() {
		CenterAlign(arts, width)
	})

	lines := strings.Split(strings.TrimSpace(output), "\n")

	for _, line := range lines {
		if len(line) > width {
			t.Errorf("Centered line exceeds width")
		}
	}
}

func TestJustifyAlign(t *testing.T) {
	arts := [][]string{
		{"A", "A", "A", "A", "A", "A", "A", "A"},
		{"B", "B", "B", "B", "B", "B", "B", "B"},
		{"C", "C", "C", "C", "C", "C", "C", "C"},
	}

	width := 30

	output := captureOutput(func() {
		JustifyAlign(arts, width)
	})

	lines := strings.Split(strings.TrimSpace(output), "\n")

	for _, line := range lines {
		if len(line) > width {
			t.Errorf("Justified line exceeds width")
		}
	}
}