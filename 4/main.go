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

func ReadInput(r io.Reader) ([]string, error) {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)

	var result []string
	var passportStr string

	for scanner.Scan() {
		if text := scanner.Text(); len(text) > 0 {
			passportStr += " " + text
		} else {
			result = append(result, passportStr)
			passportStr = ""
		}
	}

	result = append(result, passportStr)
	return result, scanner.Err()
}

func ValidatePassport(passportStr string) (bool, error) {
	r := strings.NewReader(passportStr)
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanWords)

	var okCount int
	var hasCid bool

	for scanner.Scan() {
		re := regexp.MustCompile(`(\w{3}):(.*)`)
		text := scanner.Text()
		matches := re.FindStringSubmatch(text)
		if len(matches) > 0 {
			var k, v = matches[1], matches[2]
			var ok bool
			switch k {
			case "byr":
				ok = ValidatePassportYear(v, 1920, 2002)
			case "iyr":
				ok = ValidatePassportYear(v, 2010, 2020)
			case "eyr":
				ok = ValidatePassportYear(v, 2020, 2030)
			case "hgt":
				ok = ValidatePassportHeight(v)
			case "hcl":
				ok, _ = regexp.MatchString(`#[[:xdigit:]]{6}`, v)
			case "ecl":
				ok, _ = regexp.MatchString(`amb|blu|brn|gry|grn|hzl|oth`, v)
			case "pid":
				ok, _ = regexp.MatchString(`\d{9}`, v)
			case "cid":
				ok = true
				hasCid = true
			}

			if ok {
				okCount++
			} else {
				break
			}
		}
	}

	if hasCid {
		return okCount >= 8, scanner.Err()
	} else {
		return okCount >= 7, scanner.Err()
	}
}

func ValidatePassportYear(v string, min int, max int) (matched bool) {
	re := regexp.MustCompile(`(\d{4})`)
	matches := re.FindStringSubmatch(v)
	if len(matches) == 0 {
		return false
	}
	n, _ := strconv.Atoi(matches[1])
	return n >= min && n <= max
}

func ValidatePassportHeight(v string) (matched bool) {
	re := regexp.MustCompile(`(\d+)(in|cm)`)
	matches := re.FindStringSubmatch(v)
	if len(matches) == 0 {
		return false
	}
	n, _ := strconv.Atoi(matches[1])
	unit := matches[2]
	switch unit {
	case "in":
		return n >= 59 && n <= 76
	case "cm":
		return n >= 150 && n <= 193
	default:
		panic("An unexpected unit of length matched the regex.")
	}
}

func main() {
	file, _ := os.Open("4/input.txt")
	input, _ := ReadInput(file)

	var count = 0
	for _, el := range input {
		if validation, _ := ValidatePassport(el); validation {
			count++
		}
	}
	fmt.Println(count)
}
