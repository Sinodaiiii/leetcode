package hard

func candy260210(ratings []int) int {
	preMax := 1
	preIndex := 0
	curr := 1
	ans := 1
	for i := 1; i < len(ratings); i++ {
		if ratings[i] > ratings[i-1] {
			curr += 1
			preMax = curr
			preIndex = i
			ans += curr
		} else if ratings[i] == ratings[i-1] {
			curr = 1
			preMax = curr
			preIndex = i
			ans += curr
		} else {
			curr = 1
			if i-preIndex+1 <= preMax {
				ans += i - preIndex
			} else {
				ans += i - preIndex + 1
			}
		}
		// fmt.Println(i, curr, ans)
	}
	return ans
}
