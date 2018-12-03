package day3_test

import (
	"bytes"
	"os"
	"testing"

	"github.com/JamesHageman/advent-of-code-2018/day3"
)

func TestPart1(t *testing.T) {
	test := `#1 @ 1,3: 4x4
#2 @ 3,1: 4x4
#3 @ 5,5: 2x2
`
	if x := day3.Part1(bytes.NewBufferString(test)); x != 4 {
		t.Fatalf("expected 4, got %d", x)
	}

	file, err := os.Open("./day3.txt")
	if err != nil {
		t.Fatal(err)
	}

	if res := day3.Part1(file); res != 116491 {
		t.Fatalf("wrong answer")
	}
}

func TestPart2(t *testing.T) {
}
