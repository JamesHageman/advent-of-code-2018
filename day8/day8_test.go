package day8_test

import (
	"bytes"
	"os"
	"testing"

	"github.com/JamesHageman/advent-of-code-2018/day8"
)

func TestPart1(t *testing.T) {
	x := day8.Part1(bytes.NewBufferString("2 3 0 3 10 11 12 1 1 0 1 99 2 1 1 2"))
	if x != 138 {
		t.Errorf("expected x = 138, got %d", x)
	}

	file, err := os.Open("./day8.txt")
	if err != nil {
		t.Error(err)
	}
	x = day8.Part1(file)
	t.Logf("Part 1: %d", x)
}

func TestPart2(t *testing.T) {
	x := day8.Part2(bytes.NewBufferString("2 3 0 3 10 11 12 1 1 0 1 99 2 1 1 2"))
	if x != 138 {
		t.Errorf("expected x = 138, got %d", x)
	}

	file, err := os.Open("./day8.txt")
	if err != nil {
		t.Error(err)
	}
	x = day8.Part2(file)
	t.Logf("Part 1: %d", x)
}
