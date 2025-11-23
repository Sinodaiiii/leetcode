package medium

func maxSumDivThree(nums []int) int {
	mod := [3]int{}
	for _, num := range nums {
		curr := mod
		for i := 0; i < 3; i++ {
			curr[(num+mod[i])%3] = max(curr[(num+mod[i])%3], num+mod[i])
		}
		// fmt.Println(curr)
		mod = curr
	}
	return mod[0]
}
