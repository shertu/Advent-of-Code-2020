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

type Passport struct {
	byr, iyr, eyr, hgt, hcl, ecl, pid, cid string
}

func ReadInput(r io.Reader) ([]map[string]string, error) {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)

	var result []map[string]string
	var passport = make(map[string]string, 0)

	for scanner.Scan() {
		// Check if there is an empty line to indicate a new passport.
		if text := scanner.Text(); len(text) > 0 {
			kvs, err := ReadInputLine(strings.NewReader(text))
			if err != nil {
				return result, err
			}

			// merge passport value maps
			for k, v := range kvs {
				passport[k] = v
			}
		} else {
			result = append(result, passport)
			passport = make(map[string]string, 0)
		}
	}

	//result = append(result, passport) // append the last passport
	return result, scanner.Err()
}

func ReadInputLine(r io.Reader) (map[string]string, error) {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanWords)

	var result = make(map[string]string, 0)

	for scanner.Scan() {
		k, v, err := StringConvertPassportKeyValuePair(scanner.Text())
		if err != nil {
			return result, err
		}
		result[k] = v
	}

	return result, scanner.Err()
}

func StringConvertPassportKeyValuePair(s string) (string, string, error) {
	re := regexp.MustCompile(`(\w{3}):(.*)`)
	matches := re.FindStringSubmatch(s)
	if len(matches) == 0 {
		return "", "", fmt.Errorf("\"%v\" does not match the passport key value regex", s)
	}
	return matches[1], matches[2], nil
}

func ConvertMapToPassport(m map[string]string) (Passport, bool) {
	var passport Passport

	if val, ok := m["byr"]; ok {
		passport.byr = val
	} else {
		return passport, false
	}

	if val, ok := m["iyr"]; ok {
		passport.iyr = val
	} else {
		return passport, false
	}

	if val, ok := m["eyr"]; ok {
		passport.eyr = val
	} else {
		return passport, false
	}

	if val, ok := m["hgt"]; ok {
		passport.hgt = val
	} else {
		return passport, false
	}

	if val, ok := m["hcl"]; ok {
		passport.hcl = val
	} else {
		return passport, false
	}

	if val, ok := m["ecl"]; ok {
		passport.ecl = val
	} else {
		return passport, false
	}

	if val, ok := m["pid"]; ok {
		passport.pid = val
	} else {
		return passport, false
	}

	if val, ok := m["cid"]; ok {
		passport.cid = val
	} else {
		//return passport, false
	}

	return passport, true
}

func IsPassportValuesValid(p Passport) bool {
	okHCL, _ := regexp.MatchString(`#[[:xdigit:]]{6}`, p.hcl)
	okECL, _ := regexp.MatchString(`amb|blu|brn|gry|grn|hzl|oth`, p.ecl)
	okPID, _ := regexp.MatchString(`\d{9}`, p.pid)

	return IsPassportKVValidYear(p.byr, 1920, 2002) &&
		IsPassportKVValidYear(p.iyr, 2010, 2020) &&
		IsPassportKVValidYear(p.eyr, 2020, 2030) &&
		IsPassportKVValidHeight(p.hgt) &&
		okHCL &&
		okECL &&
		okPID
}

func IsPassportKVValidYear(v string, min int, max int) (matched bool) {
	re := regexp.MustCompile(`(\d{4})`)
	matches := re.FindStringSubmatch(v)
	if len(matches) == 0 {
		return false
	}
	n, _ := strconv.Atoi(matches[1])
	return n >= min && n <= max
}

func IsPassportKVValidHeight(v string) (matched bool) {
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

	var passports []Passport
	for _, mp := range input {
		if p, ok := ConvertMapToPassport(mp); ok {
			passports = append(passports, p)
		}
	}

	var passportsWithValidValues []Passport
	for _, p := range passports {
		if ok := IsPassportValuesValid(p); ok {
			passportsWithValidValues = append(passportsWithValidValues, p)
		}
	}

	fmt.Println("part one", len(passports), "part two", len(passportsWithValidValues))
}
