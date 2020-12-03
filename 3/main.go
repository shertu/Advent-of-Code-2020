package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func readInput(r io.Reader) ([]string, error) {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)
	var result []string
	for scanner.Scan() {
		x := scanner.Text()
		result = append(result, x)
	}
	return result, scanner.Err()
}

func isTobogganEncounterTree(input []string, right int, down int) bool {
	line := input[down]
	index := right % len(line)
	rune := line[index]
	return rune == '#'
}

func countTreeEncounters(input []string, rightDelta int, downDelta int) int {
	count := 0
	rightPos := 0
	for downPos := 0; downPos < len(input); downPos += downDelta {
		if isTobogganEncounterTree(input, rightPos, downPos) {
			count++
		}
		rightPos += rightDelta
	}
	return count
}

func main() {
	file, _ := os.Open("input.txt")
	input, _ := readInput(file)

	counts := []int{
		countTreeEncounters(input, 1, 1),
		countTreeEncounters(input, 3, 1),
		countTreeEncounters(input, 5, 1),
		countTreeEncounters(input, 7, 1),
		countTreeEncounters(input, 1, 2),
	}
	
	fmt.Println(counts)
}
