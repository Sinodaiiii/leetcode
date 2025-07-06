package hot100

func topKFrequent(nums []int, k int) []int {
	list := make([]*[2]int, 0)
	dict := map[int]*[2]int{}
	n := len(nums)
	for i := 0; i < n; i++ {
		if dict[nums[i]] == nil {
			dict[nums[i]] = &[2]int{nums[i], 1}
			list = append(list, dict[nums[i]])
		} else {
			dict[nums[i]][1] += 1
		}
	}
	var quickSort func(l, r int)
	quickSort = func(l, r int) {
		if l >= r {
			return
		}
		less, curr := l, l
		for curr < r {
			if list[curr][1] > list[r][1] {
				list[curr], list[less] = list[less], list[curr]
				less += 1
			}
			curr += 1
		}
		list[less], list[r] = list[r], list[less]
		quickSort(l, less-1)
		quickSort(less+1, r)
	}
	quickSort(0, len(list)-1)
	ans := make([]int, k)
	for i := 0; i < k; i++ {
		ans[i] = list[i][0]
	}
	return ans
}
