package main

import (
	"strconv"
	"strings"
)

type pos struct {
	y int64
	x int64
}

func abs(x int64) int64 {
	if x < 0 {
		return -x
	}
	return x
}

func areAdjacent(h, t pos) bool {
	return abs(h.y-t.y) <= 1 && abs(h.x-t.x) <= 1
}

func move(p pos, dir string) (toPos pos) {
	toPos = p
	switch dir {
	case "U":
		toPos.y = p.y - 1
	case "D":
		toPos.y = p.y + 1
	case "L":
		toPos.x = p.x - 1
	case "R":
		toPos.x = p.x + 1
	}
	return
}

func follow(h pos, t pos) pos {

	xDiff := abs(h.x - t.x)
	yDiff := abs(h.y - t.y)

	if xDiff > yDiff {
		if h.x < t.x {
			return pos{h.y, h.x + 1} // left
		} else {
			return pos{h.y, h.x - 1} // right
		}
	} else if yDiff > xDiff {
		if h.y < t.y {
			return pos{h.y + 1, h.x} // up
		} else {
			return pos{h.y - 1, h.x} // down
		}
	} else if xDiff == yDiff { // diff always 2
		if h.x < t.x && h.y < t.y { // up left
			return pos{h.y + 1, h.x + 1}
		} else if h.x > t.x && h.y < t.y { // up right
			return pos{h.y + 1, h.x - 1}
		} else if h.x > t.x && h.y > t.y { // down right
			return pos{h.y - 1, h.x - 1}
		} else if h.x < t.x && h.y > t.y { // down left
			return pos{h.y - 1, h.x + 1}
		}
	}
	panic("oh noes")
}

func parseMove(m string) (string, int64) {
	s := strings.Split(m, " ")
	n, _ := strconv.ParseInt(s[1], 10, 64)
	return s[0], n
}
