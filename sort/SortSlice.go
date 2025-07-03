package sort

import "sort"

func SortSliceOfSlice(nums []int) {
	i := 1 // start index of the slice
	sort.Slice(nums[i+1:], func(a, b int) bool { return nums[i+1+a] <= nums[i+1+b] })
	// a&b is the index of new slice[0,1,2...]
	//which need to be offset with i in the original slice nums[i+1+0,1,2...]
}
