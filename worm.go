package main

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

func getChange(x int64) int64 {
	if x == 0 {
		return 0
	}
	if x < 0 {
		return -1
	}
	return 1
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
	xDiff := h.x - t.x
	yDiff := h.y - t.y

	if abs(xDiff) > 1 || abs(yDiff) > 1 {
		return pos{t.y + getChange(yDiff), t.x + getChange(xDiff)}
	}
	return pos{t.y, t.x}
}
