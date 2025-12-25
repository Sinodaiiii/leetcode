package medium

import "math"

func sortReap(rooti int, minReap []int) {
	length := len(minReap)
	mini := rooti
	if rooti*2 < length && minReap[mini] > minReap[2*rooti] {
		mini = 2 * rooti
	}
	if rooti*2+1 < length && minReap[mini] > minReap[2*rooti+1] {
		mini = 2*rooti + 1
	}
	if mini != rooti {
		minReap[rooti], minReap[mini] = minReap[mini], minReap[rooti]
		sortReap(mini, minReap)
	}
}

func findKthLargest(nums []int, k int) int {
	minReap := make([]int, k+1)
	for i := range minReap {
		minReap[i] = math.MinInt32
	}
	for _, num := range nums {
		if num >= minReap[1] {
			minReap[1] = num
			sortReap(1, minReap)
		}
	}
	return minReap[1]
}

func findKthLargest2(nums []int, k int) int {
	// fmt.Println(nums, k)
	// if k == 1 {
	//     return nums[0]
	// }
	// fmt.Println(nums, k)
	if len(nums) == 1 {
		return nums[0]
	}
	pivot := nums[0]
	l, r := -1, len(nums)
	for l < r {
		for l++; nums[l] > pivot; l++ {
		}
		for r--; nums[r] < pivot; r-- {
		}
		if l < r {
			nums[l], nums[r] = nums[r], nums[l]
		}
	}
	// fmt.Println(nums, l, r)

	if r >= k-1 {
		return findKthLargest(nums[:r+1], k)
	} else {
		return findKthLargest(nums[r+1:], k-r-1)
	}
}

//func findKthLargest(nums []int, k int) int {
//	heapSize := len(nums)
//	buildMaxHeap(nums, heapSize)
//	for i := len(nums) - 1; i >= len(nums)-k+1; i-- {
//		nums[0], nums[i] = nums[i], nums[0]
//		heapSize--
//		maxHeapify(nums, 0, heapSize)
//	}
//	return nums[0]
//}
//
//func buildMaxHeap(a []int, heapSize int) {
//	for i := heapSize/2 - 1; i >= 0; i-- {
//		maxHeapify(a, i, heapSize)
//	}
//}
//
//func maxHeapify(a []int, i, heapSize int) {
//	l, r, largest := i*2+1, i*2+2, i
//	if l < heapSize && a[l] > a[largest] {
//		largest = l
//	}
//	if r < heapSize && a[r] > a[largest] {
//		largest = r
//	}
//	if largest != i {
//		a[i], a[largest] = a[largest], a[i]
//		maxHeapify(a, largest, heapSize)
//	}
//}
