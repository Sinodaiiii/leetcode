package sort

import "fmt"

func quickSort(arr []int, low, high int) {
	if low >= high {
		return
	}

	// 选择基准值（这里选择最后一个元素）
	pivot := arr[high]

	// 分区操作
	i := low
	for j := low; j < high; j++ {
		if arr[j] < pivot {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}

	// 将基准值放到正确位置
	arr[i], arr[high] = arr[high], arr[i]

	// 递归排序左右分区
	quickSort(arr, low, i-1)
	quickSort(arr, i+1, high)
}

func HoareQuickSort(nums []int) []int {
	var sortA func(l, r int)
	sortA = func(l, r int) {
		fmt.Println("sort: ", l, r, nums)
		if l >= r {
			return
		}
		pivot := nums[l]
		i, j := l-1, r+1
		for i < j {
			for i++; nums[i] < pivot; i++ {
			}
			for j--; nums[j] > pivot; j-- {
			}
			if i < j {
				nums[i], nums[j] = nums[j], nums[i]
			}
			fmt.Println(i, j, nums)
		}
		sortA(l, j)
		sortA(j+1, r)
		// 每次循环可以保证 nums[l,i) < pivot, nums(j,r] > pivot
		// 		无法确定 l 和 r 指向的值是否和 pivot 相等，可能出现 nums[l] > pivot && nums[r] < pivot
		// 所以划分要分为
		//		[l,i-1] [i,r] 左侧严格小于，右侧严格大于等于
		// 		或 [l,j] [j+1,r] 左侧严格大于等于，右侧严格小于
		// 有 pivot 存在，使 i, j 遍历过程至少会停在 pivot 处，但若 i 停留在 0 则会出现 [0,r] 循环
		// 每次 [j+1] 都会使范围缩小，所以选 j 作为分割点
	}
	sortA(0, len(nums)-1)
	return nums
}

func main() {
	arr := []int{10, 80, 30, 90, 40, 50, 70}
	fmt.Println("Original array:", arr)
	quickSort(arr, 0, len(arr)-1)
	fmt.Println("Sorted array:  ", arr)
}
