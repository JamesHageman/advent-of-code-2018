package day1_test

import (
	"bytes"
	"os"
	"testing"

	"github.com/JamesHageman/advent-of-code-2018/day1"
)

func TestPart1(t *testing.T) {
	x := day1.Part1(bytes.NewBufferString("6 -2 -1"))
	if x != 3 {
		t.Errorf("expected x = 3, got %d", x)
	}

	file, err := os.Open("./day1.txt")
	if err != nil {
		t.Error(err)
	}
	x = day1.Part1(file)
	t.Logf("Part 1: %d", x)
}

func TestPart2(t *testing.T) {
	x := day1.Part2(bytes.NewBufferString("3 3 4 -2 -4"))
	if x != 10 {
		t.Errorf("expected x = 10, got %d", x)
	}
	x = day1.Part2(bytes.NewBufferString("1 -1"))
	if x != 0 {
		t.Errorf("expected x = 0, got %d", x)
	}

	file, err := os.Open("./day1.txt")
	if err != nil {
		t.Error(err)
	}
	x = day1.Part2(file)
	t.Logf("Part 2: %d", x)
}
