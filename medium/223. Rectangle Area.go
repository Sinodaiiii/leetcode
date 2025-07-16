package medium

func computeArea(ax1 int, ay1 int, ax2 int, ay2 int, bx1 int, by1 int, bx2 int, by2 int) int {
	x, y := 0, 0
	if ax1 <= bx1 {
		if ax2 >= bx2 {
			x = bx2 - bx1
		} else if bx1 < ax2 {
			x = ax2 - bx1
		}
	} else {
		if bx2 >= ax2 {
			x = ax2 - ax1
		} else if ax1 < bx2 {
			x = bx2 - ax1
		}
	}

	if ay1 <= by1 {
		if ay2 >= by2 {
			y = by2 - by1
		} else if by1 < ay2 {
			y = ay2 - by1
		}
	} else {
		if by2 >= ay2 {
			y = ay2 - ay1
		} else if ay1 < by2 {
			y = by2 - ay1
		}
	}

	area1 := (ax2 - ax1) * (ay2 - ay1)
	area2 := (bx2 - bx1) * (by2 - by1)
	return area1 + area2 - x*y
}
