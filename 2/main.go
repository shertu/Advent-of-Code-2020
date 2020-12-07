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

func ReadInput(r io.Reader) ([]Password, error) {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)
	var passwords []Password
	for scanner.Scan() {
		password, err := StringConvertPassword(scanner.Text())
		if err != nil {
			return passwords, err
		}
		passwords = append(passwords, *password)
	}
	return passwords, scanner.Err()
}

func StringConvertPassword(s string) (*Password, error) {
	re := regexp.MustCompile(`^(\d+)-(\d+) (\w): (\w*)$`)
	matches := re.FindStringSubmatch(s)

	if len(matches) == 0 {
		return nil, fmt.Errorf("\"%v\" does not match the password regex", s)
	}

	min, _ := strconv.Atoi(matches[1])
	max, _ := strconv.Atoi(matches[2])

	result := Password{
		min:    min,
		max:    max,
		substr: matches[3],
		str:    matches[4],
	}

	return &result, nil
}

func ValidatePasswordPartOne(en Password) bool {
	substrCount := strings.Count(en.str, en.substr)
	return substrCount >= en.min && substrCount <= en.max
}

func ValidatePasswordPartTwo(en Password) bool {
	var a, b = en.str[en.min-1], en.str[en.max-1]
	return (string(a) == en.substr) != (string(b) == en.substr) // xor between booleans is not equals
}

func main() {
	file, _ := os.Open("2/input.txt")
	input, _ := ReadInput(file)

	var countPartOne, countPartTwo int
	for _, el := range input {
		if ValidatePasswordPartOne(el) {
			countPartOne++
		}

		if ValidatePasswordPartTwo(el) {
			countPartTwo++
		}
	}
	fmt.Println("part one", countPartOne, "part two", countPartTwo)
}
