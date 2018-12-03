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
	}

	fmt.Printf("rects: %d\n", len(ret))
	return ret
}

func computeClaims(rects map[uint]rectangle) map[[2]uint][]uint {
	points := make(map[[2]uint][]uint)

	for rid, r := range rects {
		for x := r.x; x < r.x+r.width; x++ {
			for y := r.y; y < r.y+r.height; y++ {
				pt := [2]uint{x, y}
				claims := points[pt]
				claims = append(claims, rid)
				points[pt] = claims
			}
		}
	}

	return points
}

func Part1(in io.Reader) uint {
	var sum uint

	rects := scanRects(in)
	points := computeClaims(rects)

	for _, claims := range points {
		if len(claims) > 1 {
			sum++
		}
	}
	return sum
}

func Part2(in io.Reader) uint {
	noOverlaps := make(map[uint]bool)

	rects := scanRects(in)
	var points = computeClaims(rects)

	for rid := range rects {
		noOverlaps[rid] = true
	}

	for _, claims := range points {
		if len(claims) > 1 {
			for rid := range claims {
				delete(noOverlaps, uint(rid))
			}
		}
	}

	fmt.Printf("%#v\n", noOverlaps)

	for rid := range noOverlaps {
		return rid
	}
	return 0
}
