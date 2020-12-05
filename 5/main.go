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

func ConvertToSeatId(code string) int {
	var id = 0
	for i, r := range code {
		switch r {
		case 'F':
		case 'B':
			id++
		case 'L':
		case 'R':
			id++
		default:
			panic(fmt.Sprintf("%v is an invalid rune", r))
		}

		if i < len(code)-1 {
			id <<= 1
		}
	}

	return id
}

func main() {
	file, _ := os.Open("5/input.txt")
	input, _ := ReadInput(file)

	markedSeatCollection := make([]bool, 256*8)

	for _, element := range input {
		seatId := ConvertToSeatId(element)
		markedSeatCollection[seatId] = true
	}

	for i := 1; i < len(markedSeatCollection)-1; i++ {
		if markedSeatCollection[i-1] && !markedSeatCollection[i] && markedSeatCollection[i+1] {
			fmt.Println("Your seat's id is", i)
		}
	}
}
