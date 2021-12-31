package adventUtil

import (
	"os"
	"testing"
)

// It reads a file of integers delimited by commas.
func TestReadIntegersCommaf(t *testing.T) {
	// Test content to read.
	testContent := "1,2,3\n4,5,6\n7,8,9"

	// Create a test file.
	f, _ := os.CreateTemp("/tmp", "aocutil_io_test")
	defer os.Remove(f.Name())
	f.Write([]byte(testContent))
	f.Close()

	got := ReadIntegersCommaf(f.Name())
	expected := [][]int{
		[]int{1, 2, 3},
		[]int{4, 5, 6},
		[]int{7, 8, 9},
	}

	for i, row := range expected {
		for j, _ := range row {
			if got[i][j] != expected[i][j] {
				t.Errorf("Got: %v Expected: %v\n", got, expected)
			}
		}
	}
}

// It reads a file of integers delimited by spaces.
func TestReadIntegersSpacef(t *testing.T) {
	// Test content to read.
	testContent := "1 2 3\n4 5 6\n7 8 9"

	// Create a test file.
	f, _ := os.CreateTemp("/tmp", "aocutil_io_test")
	defer os.Remove(f.Name())
	f.Write([]byte(testContent))
	f.Close()

	got := ReadIntegersSpacef(f.Name())
	expected := [][]int{
		[]int{1, 2, 3},
		[]int{4, 5, 6},
		[]int{7, 8, 9},
	}

	for i, row := range expected {
		for j, _ := range row {
			if got[i][j] != expected[i][j] {
				t.Errorf("Got: %v Expected: %v\n", got, expected)
			}
		}
	}
}

// It reads a file of strings and returns regex matches.
func TestReadPatternf(t *testing.T) {
	// Test content to read.
	testContent := "abc abc|def\nghi ghi|jkl\nmno mno|pqr"
	pattern := `\w+`

	// Create a test file.
	f, _ := os.CreateTemp("/tmp", "aocutil_io_test")
	defer os.Remove(f.Name())
	f.Write([]byte(testContent))
	f.Close()

	got := ReadPatternf(f.Name(), pattern)
	expected := [][]string{
		[]string{"abc", "abc", "def"},
		[]string{"ghi", "ghi", "jkl"},
		[]string{"mno", "mno", "pqr"},
	}

	for i, row := range expected {
		for j, _ := range row {
			if got[i][j] != expected[i][j] {
				t.Errorf("Got: %v Expected: %v\n", got, expected)
			}
		}
	}
}

// It reads a file of strings and returns regex grouped matches.
func TestReadSubPatternf(t *testing.T) {
	// Test content to read.
	testContent := "abc abc|def\nghi ghi|jkl\nmno mno|pqr"
	pattern := `([\w ]+)`

	// Create a test file.
	f, _ := os.CreateTemp("/tmp", "aocutil_io_test")
	defer os.Remove(f.Name())
	f.Write([]byte(testContent))
	f.Close()

	got := ReadPatternf(f.Name(), pattern)
	expected := [][]string{
		[]string{"abc abc", "def"},
		[]string{"ghi ghi", "jkl"},
		[]string{"mno mno", "pqr"},
	}

	for i, row := range expected {
		for j, _ := range row {
			if got[i][j] != expected[i][j] {
				t.Errorf("Got: %v Expected: %v\n", got, expected)
			}
		}
	}
}

// It reads a file of integers delimited by a character.
func TestreadIntegersf(t *testing.T) {
	// Test content to read.
	testContent := "1 2 3\n4 5 6\n7 8 9"

	// Create a test file.
	f, _ := os.CreateTemp("/tmp", "aocutil_io_test")
	defer os.Remove(f.Name())
	f.Write([]byte(testContent))
	f.Close()

	got := readIntegersf(f.Name(), " ")
	expected := [][]int{
		[]int{1, 2, 3},
		[]int{4, 5, 6},
		[]int{7, 8, 9},
	}

	for i, row := range expected {
		for j, _ := range row {
			if got[i][j] != expected[i][j] {
				t.Errorf("Got: %v Expected: %v\n", got, expected)
			}
		}
	}
}

// It reads a file of integers delimited by a character.
func TestparseIntegerLineBySeperator(t *testing.T) {
	testLine := "1,2,3,4,5"
	testSeperator := ","

	got := parseIntegerLineBySeperator(testLine, testSeperator)
	expected := []int{1, 2, 3, 4, 5}

	for i, _ := range expected {
		if got[i] != expected[i] {
			t.Errorf("Got: %v Expected: %v\n", got, expected)
		}
	}
}
