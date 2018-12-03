package day3

import (
	"fmt"
	"io"
)

type rectangle struct {
	x, y, width, height uint
}

func scanRects(in io.Reader) map[uint]rectangle {
	ret := make(map[uint]rectangle)

	for {
		var id uint
		r := rectangle{}
		_, err := fmt.Fscanf(
			in, "#%d @ %d,%d: %dx%d\n",
			&id, &r.x, &r.y, &r.width, &r.height,
		)
		if err != nil {
			break
		}
		ret[id] = r
		//_, _ = fmt.Fscanf(in, "%s") // eat the rest of the line, ignore error
	}

	fmt.Printf("rects: %d\n", len(ret))
	return ret
}

func Part1(in io.Reader) uint {
	var sum uint
	points := make(map[[2]uint]uint)

	rects := scanRects(in)

	for _, r := range rects {
		for x := r.x; x < r.x+r.width; x++ {
			for y := r.y; y < r.y+r.height; y++ {
				points[[2]uint{x, y}]++
			}
		}
	}

	for _, layers := range points {
		if layers > 1 {
			sum++
		}
	}
	return sum
}

func Part2(in io.Reader) {
}
