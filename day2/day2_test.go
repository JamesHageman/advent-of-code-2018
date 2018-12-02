package day2_test

import (
	"bytes"
	"os"
	"testing"

	"github.com/JamesHageman/advent-of-code-2018/day2"
)

func TestPart1(t *testing.T) {
	test := "ababa"
	if x := day2.Part1(bytes.NewBufferString(test)); x != 1 {
		t.Fatalf("expected 1, got %d", x)
	}

	test = "ababa bababc"
	if x := day2.Part1(bytes.NewBufferString(test)); x != 4 {
		t.Fatalf("expected 4, got %d", x)
	}

	test = "ababa bababc aabb"
	if x := day2.Part1(bytes.NewBufferString(test)); x != 6 {
		t.Fatalf("expected 6, got %d", x)
	}

	test = "aaabbbccdd aaabbb cccddd"
	if x := day2.Part1(bytes.NewBufferString(test)); x != 3 {
		t.Fatalf("expected 3, got %d", x)
	}

	test = `abcdef
bababc
abbcde
abcccd
aabcdd
abcdee
ababab`
	if x := day2.Part1(bytes.NewBufferString(test)); x != 12 {
		t.Fatalf("expected 12, got %d", x)
	}

	file, err := os.Open("./day2.txt")
	if err != nil {
		t.Fatal(err)
	}
	res := day2.Part1(file)
	if res != 5368 {
		t.Fatalf("expected 5368, got %d", res)
	}
}

func TestPart2(t *testing.T) {
}
