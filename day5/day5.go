package day5

import (
	"bufio"
	"container/list"
	"fmt"
	"io"
	"strings"
)

func reacts(a, b rune) bool {
	return strings.ToLower(string(a)) == strings.ToLower(string(b)) && a != b
}

func debug(polymer *list.List) {
	for curr := polymer.Front(); curr != nil; curr = curr.Next() {
		fmt.Printf("%s", string(curr.Value.(rune)))
	}
	fmt.Printf("\n")
}

func reduce(polymer *list.List) {
	curr := polymer.Front()
	for curr != nil && curr.Next() != nil {
		a := curr.Value.(rune)
		b := curr.Next().Value.(rune)

		if !reacts(a, b) {
			curr = curr.Next()
			continue
		}

		if curr == polymer.Front() {
			oldCurr := curr
			curr = curr.Next().Next()
			polymer.Remove(oldCurr.Next())
			polymer.Remove(oldCurr)
		} else {
			curr = curr.Prev()
			polymer.Remove(curr.Next().Next())
			polymer.Remove(curr.Next())
		}
	}
}

func Part1(in io.Reader) int {
	scanner := bufio.NewScanner(in)
	scanner.Scan()

	polymer := list.New()

	for _, c := range scanner.Text() {
		polymer.PushBack(rune(c))
	}

	reduce(polymer)

	return polymer.Len()
}

func Part2(in io.Reader) {
}
