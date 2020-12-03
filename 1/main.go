package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
)

func readInput(r io.Reader) ([]int, error) {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)
	var result []int
	for scanner.Scan() {
		x, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return result, err
		}
		result = append(result, x)
	}
	return result, scanner.Err()
}

func findComponents(input []int, componentSumTarget int, componentToFindCount int) []int {
	if componentToFindCount < 0 {
		panic("The number of components must be positive.")
	}

	if componentToFindCount > len(input) {
		return nil
	}

	if componentToFindCount == 0 && componentSumTarget == 0 {
		return []int{}
	}

	for i, expense := range input {
		output := []int{expense}

		// simple case
		if componentToFindCount == 1 {
			if expense == componentSumTarget {
				return output
			}
		}

		// complex case
		if componentToFindCount > 1 {
			remainder := componentSumTarget - expense

			// only ever need to search forward for values
			for _, result := range findComponents(input[i:], remainder, componentToFindCount-1) {
				output = append(output, result)
			}

			if len(output) >= componentToFindCount {
				return output
			}
		}
	}

	return nil
}

func main() {
	file, _ := os.Open("input.txt")
	input, _ := readInput(file)

	sort.Ints(input) // sorting inputs improves the average efficiency of the algorithm
	var components = findComponents(input, 2020, 3)

	fmt.Println(components)
}
