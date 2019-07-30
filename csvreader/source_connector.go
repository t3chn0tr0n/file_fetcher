package csvreader

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// ConnectToSource :
func ConnectToSource() (*os.File, rune, error) {
	input := bufio.NewReader(os.Stdin)

	// File Path
	fmt.Print("\nEnter the Path to file: ")
	path, _ := input.ReadString('\n')
	path = strings.TrimSpace(path)

	// Read Seperator
	fmt.Print("\nEnter Seperator, Leave blank for comma ',': ")
	separator, _, _ := input.ReadRune()
	if separator == 13 || separator == ' ' {
		separator = ','
	}

	// Trying to open file
	file, err := os.Open(path)
	return file, separator, err
}
