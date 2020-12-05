package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func ReadInput(r io.Reader) ([]string, error) {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)
	var result []string
	for scanner.Scan() {
		result = append(result, scanner.Text())
	}
	return result, scanner.Err()
}

func IsTobogganEncounterTree(input []string, right int, down int) bool {
	line := input[down]
	index := right % len(line)
	return line[index] == '#'
}

func CountTreeEncounters(input []string, rightDelta int, downDelta int) int {
	count := 0
	for downPos, rightPos := 0, 0; downPos < len(input); downPos, rightPos = downPos+downDelta, rightPos+rightDelta {
		if IsTobogganEncounterTree(input, rightPos, downPos) {
			count++
		}
	}
	return count
}

func main() {
	file, _ := os.Open("3/input.txt")
	input, _ := ReadInput(file)

	counts := []int{
		CountTreeEncounters(input, 1, 1),
		CountTreeEncounters(input, 3, 1),
		CountTreeEncounters(input, 5, 1),
		CountTreeEncounters(input, 7, 1),
		CountTreeEncounters(input, 1, 2),
	}

	fmt.Println(counts)
}
