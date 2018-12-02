package day2

import (
	"bufio"
	"io"
)

func characterFrequencies(s string) (ret map[rune]uint) {
	ret = make(map[rune]uint)
	for _, c := range s {
		ret[c]++
	}
	return ret
}

func Part1(in io.Reader) uint {
	freqFreqs := make(map[uint]uint)

	scanner := bufio.NewScanner(in)
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		freqs := characterFrequencies(scanner.Text())
		uniqFreqs := make(map[uint]bool)
		for _, f := range freqs {
			uniqFreqs[f] = true
		}
		for f := range uniqFreqs {
			freqFreqs[f]++
		}
	}

	twos, _ := freqFreqs[2]
	threes, _ := freqFreqs[3]
	return twos * threes
}

func Part2() {
}
