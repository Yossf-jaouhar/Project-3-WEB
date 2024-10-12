package funcs

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// ChangeToArt reads an ASCII art file for the given banner name and returns the art as a string.
func ChangeToArt(banner string, usertext string) (string, error) {
	// Open the file for reading
	MyFile, err := os.Open("Banners/" + banner + ".txt") // Ensure the path is correct
	if err != nil {
		// Handle the error, e.g., file not found
		return "", fmt.Errorf("error opening file: %w", err) // Return the error with context
	}
	// Ensure the file gets closed when we're done
	defer MyFile.Close()

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(MyFile)

	// Slice to hold the lines from the file
	var lines []string

	// Read the file line by line
	for scanner.Scan() {
		line := scanner.Text()      // Get the current line
		lines = append(lines, line) // Append the line to the slice
	}

	// Check for errors during scanning
	if err := scanner.Err(); err != nil {
		return "", fmt.Errorf("error reading file: %w", err) // Return the error with context
	}

	// Ensure there are enough lines to process
	if len(lines) < 9 {
		return "", fmt.Errorf("the file must contain at least 9 lines of ASCII art") // Check if there's enough ASCII art
	}

	result := TheArt(lines, usertext)
	return result, nil
}

func TheArt(lines []string, text string) string {
	// Validate the input: Check for printable ASCII characters and newline/carriage return.
	for _, char := range text {
		if !(char >= 32 && char <= 126 || char == '\n' || char == '\r') {
			return "invalid character: " + string(char) + "\n"
		}
	}

	// Split input by CRLF (Windows style) newlines.
	data := strings.Split(text, "\r\n")
	result := ""

	// Skip the first line from the banner file (an empty line).
	if len(lines) > 0 {
		lines = lines[1:] // Skip the first line if it exists
	}

	// Process each line of input to generate ASCII art.
	for _, line := range data {
		if line == "" {
			result += "\n"
			continue
		}
		// For each input line, convert each character to ASCII art.
		for i := 0; i < 8; i++ {
			for j := 0; j < len(line); j++ {
				// Ensure character is within bounds
				if line[j] < 32 || line[j] > 126 {
					continue // Skip invalid characters
				}
				result += lines[((int(line[j])-32)*9)+i] // Convert ASCII to art
			}
			result += "\n" // New line after each row of art
		}
	}

	return result
}
