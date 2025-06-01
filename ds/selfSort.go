package ds

func QuickSort(numbers []int, start, end int) {
	if start >= end {
		return
	}
	pivot := numbers[start]
	l, r := start, end
	for l < r {
		for l < r && numbers[r] > pivot {
			r--
		}
		if l < r {
			numbers[l] = numbers[r]
			l++
		}
		for l < r && numbers[l] < pivot {
			l++
		}
		if l < r {
			numbers[r] = numbers[l]
			r--
		}
	}
	numbers[l] = pivot
	QuickSort(numbers, start, l-1)
	QuickSort(numbers, l+1, end)
}

func HeapSort(numbers []int) {
	var adjust func(int)
	n := len(numbers)
	adjust = func(root int) {
		exchange := root
		if 2*root < n {
			if numbers[2*root] > numbers[exchange] {
				exchange = 2 * root
			}
			if 2*root+1 < n {
				if numbers[2*root+1] > numbers[exchange] {
					exchange = 2*root + 1
				}
			}
		}
		if exchange != root {
			numbers[exchange], numbers[root] = numbers[root], numbers[exchange]
			adjust(exchange)
		}
	}
	for i := n / 2; i >= 0; i-- {
		adjust(i)
	}
}

//func HeapSort(arr []int)  {
//    arrLen := len(arr)
//    for i := (arrLen-2)/2; i >= 0; i-- {
//        arrJustDown(arr, i, arrLen)
//    }
//    end := arrLen - 1
//    for end != 0 {
//        arr[0], arr[end] = arr[end], arr[0]
//        arrJustDown(arr, 0, end)
//        end--
//    }
//    fmt.Println(arr)
//}
//func arrJustDown(arr []int, root, n int) {
//    parent := root
//    child := parent * 2 + 1
//    for child < n {
//        if child + 1 < n && arr[child + 1] > arr[child] {
//            child++
//        }
//        if arr[child] > arr[parent] {
//            arr[child], arr[parent] = arr[parent], arr[child]
//            parent = child
//            child = parent * 2 + 1
//        } else {
//            break
//        }
//    }
//}
