package funcs

import (
	"strings"
)

func TreatData(lines []string, input string) string {
	// Validate the input: Check for printable ASCII characters and newline/carriage return.
	for _, char := range input {
		if !(char >= 32 && char <= 126 || char == '\n' || char == '\r') {
			return "invalid character: " + string(char) + "\n"
		}
	}

	// Split input by CRLF (Windows style) newlines.
	data := strings.Split(input, "\r\n")
	result := ""

	// Skip the first line from the banner file (an empty line).
	lines = lines[1:]
	// Process each line of input to generate ASCII art.
	for _, line := range data {
		if line == "" {
			result += "\n"
			continue
		}
		// For each input line, convert each character to ASCII art.
		for i := 0; i < 8; i++ {
			for j := 0; j < len(line); j++ {
				result += lines[((int(line[j])-32)*9)+i]
			}
			result += "\n"
		}
	}

	return result
}
