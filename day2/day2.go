package day2

import (
	"bufio"
	"bytes"
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

func union(a, b string) string {
	var buf bytes.Buffer
	for i := range a {
		if a[i] == b[i] {
			buf.WriteByte(a[i])
		}
	}
	return buf.String()
}

func Part2(in io.Reader) string {
	words := make([]string, 0)
	scanner := bufio.NewScanner(in)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		words = append(words, scanner.Text())
	}

	for _, a := range words {
		for _, b := range words {
			u := union(a, b)
			if len(u) == len(a)-1 {
				return u
			}
		}
	}
	return "not found"
}
