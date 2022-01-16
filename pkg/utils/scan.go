package utils

import (
	"bufio"
	"fmt"
	"strconv"
)

// ReadInt reads int from scanner. Scanner must be split by bufio.ScanWords
func ReadInt(scanner *bufio.Scanner) (int, error) {
	if !scanner.Scan() {
		return 0, fmt.Errorf("size of graph is not specified")
	}
	return strconv.Atoi(scanner.Text())
}
