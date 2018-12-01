package day1

import (
	"fmt"
	"io"
)

func readInts(input io.Reader) []int {
	ints := make([]int, 0)
	var x int
	for {
		_, err := fmt.Fscanf(input, "%d", &x)
		if err != nil {
			break
		}
		ints = append(ints, x)
	}

	return ints
}

func Part1(input io.Reader) int {
	var sum int
	for _, i := range readInts(input) {
		sum += i
	}
	return sum
}

func Part2(input io.Reader) int {
	var sum int
	sums := make(map[int]int)
	inputs := readInts(input)

	sums[0] = 1

	i := 0
	for {
		sum += inputs[i]
		freq, _ := sums[sum]
		freq++
		if freq == 2 {
			return sum
		}
		sums[sum] = freq
		i = (i + 1) % len(inputs)
	}
}
