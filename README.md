# ASCII Art Justify

## 📌 Description

This project extends the ASCII-art generator by adding **text alignment options**.  
It allows you to display ASCII art text aligned in different ways based on the terminal width.

---

## 🎯 Objectives

- Convert input text into ASCII art using banner files
- Support multiple alignment types:
  - `left`
  - `right`
  - `center`
  - `justify`
- Adapt output dynamically to the **terminal size**
- Ensure proper formatting even when the terminal is resized

---

## ⚙️ Features

- Alignment controlled using a flag:

```bash
--align=<type>
```

- Supported alignment types:
  - `left` → default alignment
  - `right` → aligns text to the right
  - `center` → centers text
  - `justify` → spreads words evenly across the terminal width

- Automatically detects terminal width
- Handles multi-line input using `\n`
- Works with different banner styles:
  - `standard`
  - `shadow`
  - `thinkertoy`

---

## 🚀 Usage

### Basic usage

```bash
go run . [STRING]
```

---

### With alignment

```bash
go run . --align=<type> [STRING]
```

---

### With alignment and banner

```bash
go run . --align=<type> [STRING] [BANNER]
```

---

### Example

```bash
go run . --align=right "hello world" standard
```

---

## ❗ Invalid Usage

If the flag format is incorrect, the program must return:

```bash
Usage: go run . [OPTION] [STRING] [BANNER]

Example: go run . --align=right something standard
```

---

## 📏 Terminal Behavior

- Output adjusts to terminal width
- If the terminal is resized, alignment updates accordingly
- Only text that fits within the terminal width will be processed

---

## 🧠 How It Works

1. **Input Parsing**
   - Reads command-line arguments
   - Detects alignment type and banner

2. **ASCII Conversion**
   - Converts characters into ASCII art using banner files

3. **Line Handling**
   - Supports multi-line input (`\n`)
   - Splits text into logical lines

4. **Alignment Engine**
   - Calculates text width
   - Applies spacing based on alignment type

---

## 📁 Project Structure

```text
.
├── main.go
├── go.mod
├── standard.txt
├── shadow.txt
├── thinkertoy.txt
├── README.md
└── art_justify/
    ├── ascii_art.go
    ├── align.go
    ├── terminal_size.go
    └── art_justify_test.go
```

---

## 🛠️ Requirements

- Go (Golang)
- Standard Go packages only

---

## 🧪 Testing

This project includes unit tests to verify both the ASCII-art generator and alignment functions.

### Tests Covered

#### ASCII Art Tests
- Ensures ASCII conversion does not return an empty result
- Confirms each generated ASCII character contains exactly 8 rows
- Verifies newline handling (`\n`) creates multiple output blocks correctly

#### Alignment Tests
- **Left Align**
  - Checks that words are printed in normal left alignment

- **Right Align**
  - Ensures output stays within terminal width

- **Center Align**
  - Verifies centered text does not exceed terminal size

- **Justify Align**
  - Confirms spaces are distributed evenly between words
  - Ensures justified output fits within terminal width

### Test Utility

A helper function captures terminal output during tests:

```go
captureOutput()
```

This allows printed ASCII-art output to be tested and validated automatically.

---

### Running Tests

Run all tests with:

```bash
go test ./...
```

Or run tests in verbose mode:

```bash
go test -v ./...
```

---

## ⚠️ Notes

- The alignment flag must be exactly:

```bash
--align=<type>
```

- Any variation (e.g. `-align`, `align=`, etc.) is invalid

- Program must still work with:
  - only `[STRING]`
  - `[STRING] + [BANNER]`
  - `[OPTION] + [STRING]`

---

## ✅ Summary

This project focuses on:

- CLI argument parsing
- String manipulation
- Terminal-aware formatting
- Clean and modular Go code

---

## 👨‍💻 Author

ASCII Art Justify Project
