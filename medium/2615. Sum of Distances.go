package medium

func distance260423(nums []int) []int64 {
	n := len(nums)
	left := make([]int, n)
	preIndex := map[int]int{}
	preNum := map[int]int{}
	for i := 0; i < n; i++ {
		num := nums[i]
		if index, exist := preIndex[num]; exist {
			left[i] = left[index] + preNum[num]*(i-index)
		}
		preIndex[num] = i
		preNum[num] += 1
	}
	right := make([]int, n)
	preIndex = map[int]int{}
	preNum = map[int]int{}
	for i := n - 1; i >= 0; i-- {
		num := nums[i]
		if index, exist := preIndex[num]; exist {
			right[i] = right[index] + preNum[num]*(index-i)
		}
		preIndex[num] = i
		preNum[num] += 1
	}
	arr := make([]int64, n)
	for i := range arr {
		arr[i] = int64(left[i] + right[i])
	}
	return arr
}
