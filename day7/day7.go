package day7

import (
	"fmt"
	"io"
)

type runeSet map[rune]struct{}
type dag map[rune]node

var (
	present = struct{}{}
)

type node struct {
	id      rune
	in, out runeSet
}

func parseDag(in io.Reader) dag {
	g := make(dag)
	for {
		var to, from rune
		_, err := fmt.Fscanf(in, "Step %c must be finished before step %c can begin.\n", &from, &to)
		if err != nil {
			fmt.Printf("%#v\n", err)
			break
		}
		fromNode, ok := g[from]
		if !ok {
			fromNode.id = from
			fromNode.in = make(runeSet)
			fromNode.out = make(runeSet)
		}
		toNode, ok := g[to]
		if !ok {
			toNode.id = to
			toNode.in = make(runeSet)
			toNode.out = make(runeSet)
		}
		fromNode.out[to] = present
		toNode.in[from] = present

		g[from] = fromNode
		g[to] = toNode
	}
	return g
}

func Part1(in io.Reader) string {
	g := parseDag(in)
	fmt.Printf("%#v\n", g)
	return ""
}

func Part2(in io.Reader) string {
	return ""
}
