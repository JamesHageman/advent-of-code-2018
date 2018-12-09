package day8

import (
	"fmt"
	"io"
)

type Node struct {
	children []*Node
	metadata []int
}

func scan(in io.Reader) *Node {
	var c, m int
	fmt.Fscanf(in, "%d %d", &c, &m)
	node := &Node{
		children: make([]*Node, c),
		metadata: make([]int, m),
	}

	for i := 0; i < c; i++ {
		child := scan(in)
		node.children[i] = child
	}

	for i := 0; i < m; i++ {
		var d int
		fmt.Fscanf(in, "%d", &d)
		node.metadata[i] = d
	}

	return node
}

func sumMeta(node *Node) int {
	sum := 0
	for _, child := range node.children {
		sum += sumMeta(child)
	}

	for _, d := range node.metadata {
		sum += d
	}

	return sum
}

func getValue(node *Node) int {
	if len(node.children) == 0 {
		return sumMeta(node)
	}

	sum := 0
	for _, d := range node.metadata {
		if d < 1 || d > len(node.children) {
			continue
		}
		child := node.children[d-1]
		sum += getValue(child)
	}
	return sum
}

func Part1(input io.Reader) int {
	root := scan(input)
	return sumMeta(root)
}

func Part2(input io.Reader) int {
	root := scan(input)
	return getValue(root)
}
