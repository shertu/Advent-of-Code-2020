package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type PasswordEntity struct {
	min, max int
	substr   string
	str      string
}

func ReadInput(filename string) []PasswordEntity {
	file, err := os.Open(filename)

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	var entities []PasswordEntity = []PasswordEntity{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		entities = append(entities, ParsePasswordEntity(text))
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return entities
}

func ParsePasswordEntity(text string) PasswordEntity {
	re := regexp.MustCompile(`^(\d+)-(\d+) (\w): (\w*)$`)
	submatches := re.FindStringSubmatch(text)

	var minMatch, maxMatch, letterMatch, valueMatch string = submatches[1], submatches[2], submatches[3], submatches[4]

	// min and max match
	min, _ := strconv.Atoi(minMatch)
	max, _ := strconv.Atoi(maxMatch)

	return PasswordEntity{
		min:    min,
		max:    max,
		substr: letterMatch,
		str:    valueMatch,
	}
}

func ValidatePasswordEntityAlpha(en PasswordEntity) bool {
	substrCount := strings.Count(en.str, en.substr)
	return substrCount >= en.min && substrCount <= en.max
}

func ValidatePasswordEntityBravo(en PasswordEntity) bool {
	var a, b byte = en.str[en.min-1], en.str[en.max-1]
	return (string(a) == en.substr) != (string(b) == en.substr) // xor between booleans is not equals
}

func main() {
	var filename string = "input.txt"
	entities := ReadInput(filename)

	var validPasswordCount int = 0
	for i := 0; i < len(entities); i++ {
		if ValidatePasswordEntityBravo(entities[i]) {
			validPasswordCount++
		}
	}
	fmt.Println(validPasswordCount)
}
