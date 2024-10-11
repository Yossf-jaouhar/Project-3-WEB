package funcs

import (
	"bufio"
	"fmt"
	"os"
)

// ReadB opens a banner file and returns its contents as a slice of strings, line by line.
func ReadB(banner string) ([]string, error) {
	// Construct the file path.
	filePath := "Banners/" + banner + ".txt"

	// Open the file.
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("banner file '%s' not found: %w", filePath, err)
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)

	// Read each line from the file.
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	// Check for scanner errors.
	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading banner file '%s': %w", filePath, err)
	}

	// Validate the number of lines.
	if len(lines) != 855 {
		return nil, fmt.Errorf("banner file '%s' is corrupted: expected 855 lines", filePath)
	}

	return lines, nil
}
