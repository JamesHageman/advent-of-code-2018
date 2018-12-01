package day1

import (
	"fmt"
	"io"
)

func Part1(input io.Reader) int {
	var sum, i int
	for {
		_, err := fmt.Fscanf(input, "%d", &i)
		if err != nil {
			break
		}
		sum += i
	}
	return sum
}

func Part2(input io.Reader) int {
	var sum, x int
	sums := make(map[int]int)
	inputs := make([]int, 0)

	sums[0] = 1

	var err error
	for {
		_, err = fmt.Fscanf(input, "%d", &x)
		if err != nil {
			break
		}
		inputs = append(inputs, x)
	}

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
