package medium

func largestSquareArea260117(bottomLeft [][]int, topRight [][]int) int64 {
	minWidth := 0
	for i := 0; i < len(bottomLeft); i++ {
		for j := i + 1; j < len(topRight); j++ {
			currX := 0
			currY := 0
			if bottomLeft[i][0] == bottomLeft[j][0] {
				currX = min(topRight[i][0], topRight[j][0]) - bottomLeft[i][0]
			} else if bottomLeft[i][0] < bottomLeft[j][0] && topRight[i][0] > bottomLeft[j][0] {
				currX = min(topRight[i][0]-bottomLeft[j][0], topRight[j][0]-bottomLeft[j][0])
			} else if bottomLeft[j][0] < bottomLeft[i][0] && topRight[j][0] > bottomLeft[i][0] {
				currX = min(topRight[j][0]-bottomLeft[i][0], topRight[i][0]-bottomLeft[i][0])
			}

			if bottomLeft[i][1] == bottomLeft[j][1] {
				currY = min(topRight[i][1], topRight[j][1]) - bottomLeft[i][1]
			} else if bottomLeft[i][1] < bottomLeft[j][1] && topRight[i][1] > bottomLeft[j][1] {
				currY = min(topRight[i][1]-bottomLeft[j][1], topRight[j][1]-bottomLeft[j][1])
			} else if bottomLeft[j][1] < bottomLeft[i][1] && topRight[j][1] > bottomLeft[i][1] {
				currY = min(topRight[j][1]-bottomLeft[i][1], topRight[i][1]-bottomLeft[i][1])
			}
			minWidth = max(minWidth, min(currX, currY))
		}
	}
	return int64(minWidth * minWidth)
}
