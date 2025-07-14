package hard

// can be optimized by hash table in second&third for iter
func maxPoints(points [][]int) int {
	if len(points) <= 2 {
		return len(points)
	}
	ans := 2
	for i := 0; i < len(points)-1; i++ {
		for j := i + 1; j < len(points); j++ {
			curr := 2
			if points[i][0] == points[j][0] {
				for k := j + 1; k < len(points); k++ {
					if points[k][0] == points[i][0] {
						curr += 1
					}
				}
			} else {
				t := float64(points[i][1]-points[j][1]) / float64(points[i][0]-points[j][0])
				for k := j + 1; k < len(points); k++ {
					if float64(points[i][1]-points[k][1])/float64(points[i][0]-points[k][0]) == t {
						curr += 1
					}
				}
			}
			ans = max(ans, curr)
		}
	}
	return ans
}
