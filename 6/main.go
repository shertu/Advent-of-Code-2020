package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

// Copied from https://github.com/juliangruber/go-intersect
// Hash has complexity: O(n * x) where x is a factor of hash function efficiency (between 1 and 2)
func Hash(a []rune, b []rune) []rune {
	set := make([]rune, 0)
	hash := make(map[rune]bool)

	for i := 0; i < len(a); i++ {
		el := a[i]
		hash[el] = true
	}

	for i := 0; i < len(b); i++ {
		el := b[i]
		if _, found := hash[el]; found {
			set = append(set, el)
		}
	}

	return set
}

func ReadInput(r io.Reader) ([]string, error) {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)

	var result []string
	var inputElement string

	for scanner.Scan() {
		// Check if there is an empty line to indicate a new inputElement.
		if text := scanner.Text(); len(text) > 0 {
			inputElement += " " + text
		} else {
			result = append(result, inputElement)
			inputElement = ""
		}
	}

	result = append(result, inputElement) // append the last inputElement
	return result, scanner.Err()
}

func ProcessGroup(answers string) ([]rune, error) {
	r := strings.NewReader(answers)
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanWords)

	var result []rune
	for scanner.Scan() {
		runeCollection := []rune(scanner.Text())

		if result == nil {
			result = runeCollection
		} else {
			result = Hash(result, runeCollection)
		}
	}

	return result, scanner.Err()
}

func main() {
	file, _ := os.Open("6/input.txt")
	input, _ := ReadInput(file)

	count := 0
	for _, element := range input {
		x, _ := ProcessGroup(element)
		count += len(x)
	}
	fmt.Println(count)
}
