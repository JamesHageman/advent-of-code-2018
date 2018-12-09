package day7_test

import (
	"bytes"
	"os"
	"testing"

	"github.com/JamesHageman/advent-of-code-2018/day7"
)

func TestPart1(t *testing.T) {
	test := `Step C must be finished before step A can begin.
Step C must be finished before step F can begin.
Step A must be finished before step B can begin.
Step A must be finished before step D can begin.
Step B must be finished before step E can begin.
Step D must be finished before step E can begin.
Step F must be finished before step E can begin.
`

	if x := day7.Part1(bytes.NewBufferString(test)); x != "CABDFE" {
		t.Fatalf("expected CABDFE, got %s", x)
	}

	file, err := os.Open("./day7.txt")
	if err != nil {
		t.Fatal(err)
	}

	res := day7.Part1(file)
	t.Logf("Part 1: %s", res)
	//if res != "" {
	//	t.Fatalf("wrong answer")
	//}
}

func TestPart2(t *testing.T) {
	test := `Step C must be finished before step A can begin.
Step C must be finished before step F can begin.
Step A must be finished before step B can begin.
Step A must be finished before step D can begin.
Step B must be finished before step E can begin.
Step D must be finished before step E can begin.
Step F must be finished before step E can begin.
`

	if x := day7.Part2(bytes.NewBufferString(test)); x != "CABDFE" {
		t.Fatalf("expected CABDFE, got %s", x)
	}

	file, err := os.Open("./day7.txt")
	if err != nil {
		t.Fatal(err)
	}

	res := day7.Part2(file)
	t.Logf("Part 1: %s", res)
	//if res != "" {
	//	t.Fatalf("wrong answer")
	//}
}
