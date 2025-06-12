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

func main() {
	arr := []int{10, 80, 30, 90, 40, 50, 70}
	fmt.Println("Original array:", arr)
	quickSort(arr, 0, len(arr)-1)
	fmt.Println("Sorted array:  ", arr)
}
