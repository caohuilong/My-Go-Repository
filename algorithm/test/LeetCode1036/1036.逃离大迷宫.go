package main

import "fmt"

func main() {
	blocked := [][]int{{0, 1}, {1, 0}}
	source := []int{0, 0}
	target := []int{1, 1}
	res := isEscapePossible(blocked, source, target)
	fmt.Println(res)
}

type pair struct {
	x, y int
}

var dirs = []pair{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

func isEscapePossible(blocked [][]int, source []int, target []int) bool {
	const (
		block        = -1
		valid        = 0
		found        = 1
		boundary int = 1e6
	)

	n := len(blocked)
	if n < 2 {
		return true
	}

	blockset := map[pair]bool{}
	for _, b := range blocked {
		blockset[pair{b[0], b[1]}] = true
	}

	check := func(start, finish []int) int {
		sx, sy := start[0], start[1]
		fx, fy := finish[0], finish[1]
		countDown := n * (n - 1) / 2

		q := []pair{{sx, sy}}
		vis := map[pair]bool{pair{sx, sy}: true}
		for len(q) > 0 && countDown > 0 {
			p := q[0]
			q = q[1:]
			for _, d := range dirs {
				x, y := p.x+d.x, p.y+d.y
				np := pair{x, y}
				if 0 <= x && x < boundary && 0 <= y && y < boundary && !blockset[np] && !vis[np] {
					if x == fx && y == fy {
						return found
					}
					countDown--
					vis[np] = true
					q = append(q, np)
				}
			}

		}
		if countDown > 0 {
			return block
		}
		return valid
	}

	res := check(source, target)
	return res == found || res == valid && check(target, source) != block
}
