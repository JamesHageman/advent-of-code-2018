package day9

type Marble struct {
	prev, next *Marble
	value      int
}

func starter() (ret *Marble) {
	ret = &Marble{value: 0}
	ret.prev = ret
	ret.next = ret
	return
}

func Part1(numPlayers, maxMarble int) int {
	curr := starter()
	scores := make([]int, numPlayers)
	currPlayer := 0
	maxScore := 0

	for i := 1; i <= maxMarble; i++ {
		if i%23 == 0 {
			scores[currPlayer] += i

			rm := curr
			for j := 0; j < 7; j++ {
				rm = rm.prev
			}

			curr = rm.next

			rm.next.prev = rm.prev
			rm.prev.next = rm.next

			rm.prev = nil
			rm.next = nil

			scores[currPlayer] += rm.value

			if scores[currPlayer] > maxScore {
				maxScore = scores[currPlayer]
			}
		} else {
			left := curr.next
			right := left.next

			m := &Marble{value: i}
			m.prev = left
			m.next = right

			left.next = m
			right.prev = m

			curr = m
		}
		currPlayer = (currPlayer + 1) % numPlayers
	}
	return maxScore
}

func Part2(numPlayers, maxMarble int) int {
	return 0
}
