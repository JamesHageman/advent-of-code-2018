package day5_test

import (
	"bytes"
	"os"
	"testing"

	"github.com/JamesHageman/advent-of-code-2018/day5"
)

func TestPart1(t *testing.T) {
	test := "dabAcCaCBAcCcaDA"

	if x := day5.Part1(bytes.NewBufferString(test)); x != 10 {
		t.Fatalf("expected 10, got %d", x)
	}

	file, err := os.Open("./day5.txt")
	if err != nil {
		t.Fatal(err)
	}

	res := day5.Part1(file)
	t.Logf("Part 1: %d", res)
	if res != 9202 {
		t.Fatalf("wrong answer")
	}
}

func TestPart2(t *testing.T) {
}
