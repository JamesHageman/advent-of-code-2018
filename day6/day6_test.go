package day6_test

import (
	"bytes"
	"os"
	"testing"

	"github.com/JamesHageman/advent-of-code-2018/day6"
)

func TestPart1(t *testing.T) {
	test := `1, 1
1, 6
8, 3
3, 4
5, 5
8, 9
`

	if x := day6.Part1(bytes.NewBufferString(test)); x != 17 {
		t.Fatalf("expected 17, got %d", x)
	}

	file, err := os.Open("./day6.txt")
	if err != nil {
		t.Fatal(err)
	}

	res := day6.Part1(file)
	t.Logf("Part 1: %d", res)
	//if res != 77941 {
	//	t.Fatalf("wrong answer")
	//}
}

func TestPart2(t *testing.T) {
	test := `1, 1
1, 6
8, 3
3, 4
5, 5
8, 9
`

	if x := day6.Part2(bytes.NewBufferString(test)); x != 0 {
		t.Fatalf("expected 0, got %d", x)
	}

	file, err := os.Open("./day6.txt")
	if err != nil {
		t.Fatal(err)
	}

	res := day6.Part2(file)
	t.Logf("Part 2: %d", res)
	//if res != 35289 {
	//	t.Fatalf("wrong answer")
	//}
}
