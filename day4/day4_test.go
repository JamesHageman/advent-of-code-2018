package day4_test

import (
	"bytes"
	"os"
	"testing"

	"github.com/JamesHageman/advent-of-code-2018/day4"
)

func TestPart1(t *testing.T) {
	test := `[1518-11-01 00:00] Guard #10 begins shift
[1518-11-01 00:05] falls asleep
[1518-11-01 00:25] wakes up
[1518-11-01 00:30] falls asleep
[1518-11-01 00:55] wakes up
[1518-11-01 23:58] Guard #99 begins shift
[1518-11-02 00:40] falls asleep
[1518-11-02 00:50] wakes up
[1518-11-03 00:05] Guard #10 begins shift
[1518-11-03 00:24] falls asleep
[1518-11-03 00:29] wakes up
[1518-11-04 00:02] Guard #99 begins shift
[1518-11-04 00:36] falls asleep
[1518-11-04 00:46] wakes up
[1518-11-05 00:03] Guard #99 begins shift
[1518-11-05 00:45] falls asleep
[1518-11-05 00:55] wakes up
`

	if x := day4.Part1(bytes.NewBufferString(test)); x != 240 {
		t.Fatalf("expected 240, got %d", x)
	}

	file, err := os.Open("./day4.txt")
	if err != nil {
		t.Fatal(err)
	}

	res := day4.Part1(file)
	t.Logf("Part 1: %d", res)
}

func TestPart2(t *testing.T) {
}
