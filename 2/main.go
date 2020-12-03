package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Password struct {
	min, max int
	substr   string
	str      string
}

func readInput(r io.Reader) ([]Password, error) {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)
	var result []Password
	for scanner.Scan() {
		x := strconvPassword(scanner.Text())
		result = append(result, x)
	}
	return result, scanner.Err()
}

func strconvPassword(s string) Password {
	re := regexp.MustCompile(`^(\d+)-(\d+) (\w): (\w*)$`)
	submatches := re.FindStringSubmatch(s)

	if len(submatches) == 0 {
		panic(fmt.Sprintf("\"%v\" does not match the password regex.", s))
	}

	var minMatch, maxMatch, letterMatch, valueMatch string = submatches[1], submatches[2], submatches[3], submatches[4]

	min, _ := strconv.Atoi(minMatch)
	max, _ := strconv.Atoi(maxMatch)

	return Password{
		min:    min,
		max:    max,
		substr: letterMatch,
		str:    valueMatch,
	}
}

func validatePasswordAlpha(en Password) bool {
	substrCount := strings.Count(en.str, en.substr)
	return substrCount >= en.min && substrCount <= en.max
}

func validatePasswordBravo(en Password) bool {
	var a, b byte = en.str[en.min-1], en.str[en.max-1]
	return (string(a) == en.substr) != (string(b) == en.substr) // xor between booleans is not equals
}

func main() {
	file, _ := os.Open("input.txt")
	input, _ := readInput(file)

	count := 0
	for _, en := range input {
		if validatePasswordBravo(en) {
			count++
		}
	}
	fmt.Println(count)
}
