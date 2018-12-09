package day9_test

import (
	"testing"

	"github.com/JamesHageman/advent-of-code-2018/day9"
)

func TestPart1(t *testing.T) {
	tests := [][3]int{
		[3]int{9, 25, 32},
	}

	for _, in := range tests {
		numPlayers, numMarbles, highScore := in[0], in[1], in[2]
		if res := day9.Part1(numPlayers, numPlayers); res != highScore {
			t.Errorf("numPlayers=%d numMarbles=%d. Expected %d, got %d", numPlayers, numMarbles, highScore, res)
		}
	}

	x := day9.Part1(416, 71975)
	t.Logf("Part 1: %d", x)
	//if x != 41521 {
	//	t.Fatalf("wrong answer")
	//}
}

func TestPart2(t *testing.T) {
}
