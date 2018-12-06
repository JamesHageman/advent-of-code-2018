package day6

import (
	"container/list"
	"fmt"
	"io"
)

type point_t struct {
	x, y int
}

func (p *point_t) adjacent() [4]point_t {
	x, y := p.x, p.y
	return [4]point_t{
		point_t{x + 1, y},
		point_t{x - 1, y},
		point_t{x, y + 1},
		point_t{x, y - 1},
	}
}

type bounds_t struct {
	topLeft, bottomRight point_t
}

type grid_t struct {
	bounds bounds_t
	buf    [][]int
}

func (g *grid_t) width() int {
	return g.bounds.bottomRight.x - g.bounds.topLeft.x + 1
}

func (g *grid_t) height() int {
	return g.bounds.bottomRight.y - g.bounds.topLeft.y + 1
}

func newGrid(b bounds_t) (g grid_t) {
	g.bounds = b
	g.buf = make([][]int, g.width())
	for x := range g.buf {
		g.buf[x] = make([]int, g.height())
	}
	return
}

func (g *grid_t) at(point point_t) (*int, error) {
	_x := point.x - g.bounds.topLeft.x
	_y := point.y - g.bounds.topLeft.y
	if _x < 0 || _x >= len(g.buf) || _y < 0 || _y >= len(g.buf[_x]) {
		return nil, fmt.Errorf("Point out of bounds: %#v (_x=%d, _y=%d)", point, _x, _y)
	}
	return &g.buf[_x][_y], nil
}

func scanPoints(in io.Reader) (points []point_t) {
	for {
		var point point_t
		_, err := fmt.Fscanf(in, "%d, %d\n", &point.x, &point.y)
		if err != nil {
			break
		}
		points = append(points, point)
	}
	return
}

func getBounds(points []point_t) (bounds bounds_t) {
	bounds.topLeft = points[0]
	bounds.bottomRight = points[0]

	for _, p := range points {
		if p.x < bounds.topLeft.x {
			bounds.topLeft.x = p.x
		}
		if p.x > bounds.bottomRight.x {
			bounds.bottomRight.x = p.x
		}
		if p.y < bounds.topLeft.y {
			bounds.topLeft.y = p.y
		}
		if p.y > bounds.bottomRight.y {
			bounds.bottomRight.y = p.y
		}
	}

	// extend by one
	bounds.topLeft.x--
	bounds.topLeft.y--
	bounds.bottomRight.x++
	bounds.bottomRight.y--
	return
}

const (
	shared = -1
	empty  = 0
)

func flood(grid *grid_t, points []point_t) []int {
	sizes := make([]int, len(points))
	q := list.New()

	type elem struct {
		index int
		point point_t
	}

	for i, p := range points {
		id := i + 1
		q.PushBack(elem{i, p})
		cell, err := grid.at(p)
		if err != nil {
			panic(fmt.Errorf("Unexpected bad point: %v", err))
		}
		*cell = id
		sizes[i]++
	}

	for q.Len() > 0 {
		e := q.Remove(q.Front()).(elem)
		p := e.point
		i := e.index
		id := i + 1 // positve cell indicated occupation from a point

		for _, p := range p.adjacent() {
			cell, err := grid.at(p)

			if err != nil {
				fmt.Printf("%v\n", err)
				continue
			}

			if *cell == empty {
				sizes[i]++
				*cell = id
				q.PushBack(elem{i, p})
			} else if *cell != shared {
				otherIndex := *cell - 1
				sizes[otherIndex]--
				*cell = shared
			}
		}
	}

	return sizes
}

func Part1(in io.Reader) int {
	points := scanPoints(in)
	fmt.Printf("points: %#v\n", points)
	bounds := getBounds(points)
	grid := newGrid(bounds)
	sizes := flood(&grid, points)
	fmt.Printf("sizes: %#v\n", sizes)

	maxSize := sizes[0]
	for _, size := range sizes {
		if size > maxSize {
			maxSize = size
		}
	}

	return maxSize
}

func Part2(in io.Reader) int {
	return 1
}
