package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
)

func ReadInput(r io.Reader) ([]int, error) {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanWords)

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

func MaxInt(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}

func ExtendedKnapsack(profit []int, weight []int, n int, maxW int, maxE int) [][][]int {
	//fmt.Println(profit, weight, n, maxW, maxE)
	var dx, dy, dz = n + 1, maxW + 1, maxE + 1

	// Construct multi-dimensional slice.
	var dp = make([][][]int, dx)
	for i := range dp {
		dp[i] = make([][]int, dy)
		for j := range dp[i] {
			dp[i][j] = make([]int, dz)
		}
	}

	// For each element given,
	for i := 1; i < dx; i++ {
		// For each possible weight value,
		for j := 1; j < dy; j++ {
			// For each case where the total elements are less than the constraint,
			for k := 1; k < dz; k++ {
				// To ensure that we dont go out of the array,
				if j >= weight[i-1] {
					dp[i][j][k] = MaxInt(
						dp[i-1][j][k],
						dp[i-1][j-weight[i-1]][k-1]+profit[i-1],
					)
				} else {
					dp[i][j][k] = dp[i-1][j][k]
				}
			}
		}
	}

	return dp
}

func TraceExtendedKnapsack(dp [][][]int, profit []int, weight []int, n int, maxW int, maxE int) []int {
	trace := make([]int, 0)

	for i, j, k := n, maxW, maxE; i >= 0 && j >= 0 && k >= 0; i-- {
		elem := dp[i][j][k]
		//fmt.Println(i, j, k, elem)

		if i > 0 {
			elemPrev := dp[i-1][j][k]
			if elem != elemPrev {
				trace = append(trace, profit[i-1])
				j -= weight[i-1]
				k--
			}
		}
	}

	return trace
}

func main() {
	file, _ := os.Open("1/input.txt")
	input, _ := ReadInput(file)

	sort.Ints(input) // sorting causes items to be tried from smallest to largest

	n := len(input)
	maxW := 2020
	maxE := 3

	res := ExtendedKnapsack(input, input, n, maxW, maxE)
	trace := TraceExtendedKnapsack(res, input, input, n, maxW, maxE)
	fmt.Println(trace)
}
