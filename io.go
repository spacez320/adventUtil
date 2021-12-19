package adventUtil

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

// Reads a file of lines of numbers separated by commas.
func ReadIntegersCommaf(file string) [][]int {
	return readIntegersf(file, ",")
}

// Reads a file of lines of integers separated by spaces.
func ReadIntegersSpacef(file string) [][]int {
	return readIntegersf(file, " ")
}

// Reads a file full of lines of integers with a common delimiter and returns a
// slice of slices of those integers.
func readIntegersf(file, seperator string) (ints [][]int) {
	inputf, _ := os.Open(file)
	defer inputf.Close()

	scanner := bufio.NewScanner(inputf)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		ints = append(
			ints,
			parseIntegerLineBySeperator(scanner.Text(), seperator))
	}

	return
}

// Seperates a line of delimited integers into a slice of those integers.
func parseIntegerLineBySeperator(line, seperator string) (numbers []int) {
	for _, num := range strings.SplitN(line, seperator, -1) {
		next, _ := strconv.Atoi(strings.TrimSpace(num))
		numbers = append(numbers, next)
	}

	return
}
