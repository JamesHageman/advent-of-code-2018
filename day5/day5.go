package day5

import (
	"bufio"
	"container/list"
	"io"
	"unicode"
)

func reacts(a, b rune) bool {
	return unicode.ToLower(a) == unicode.ToLower(b) && a != b
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
		polymer.PushBack(c)
	}

	reduce(polymer)

	return polymer.Len()
}

func Part2(in io.Reader) int {
	scanner := bufio.NewScanner(in)
	scanner.Scan()
	text := scanner.Text()

	minSize := len(text)

	for ignoreCh := 'a'; ignoreCh <= 'z'; ignoreCh++ {
		polymer := list.New()

		for _, c := range scanner.Text() {
			if unicode.ToLower(c) != ignoreCh {
				polymer.PushBack(c)
			}
		}

		reduce(polymer)

		if polymer.Len() < minSize {
			minSize = polymer.Len()
		}
	}

	return minSize
}
