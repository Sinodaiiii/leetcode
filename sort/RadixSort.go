package sort

import "math"

// 基数排序（LSD：从最低位开始）
func radixSort(arr []int) {
	if len(arr) <= 1 {
		return
	}

	// 找到数组中的最大值，确定最大位数
	maxVal := arr[0]
	for _, val := range arr {
		if val > maxVal {
			maxVal = val
		}
	}

	// 计算最大位数
	maxDigit := 0
	for maxVal > 0 {
		maxVal /= 10
		maxDigit++
	}

	// 用于计数排序的桶
	bucket := make([][]int, 10)

	// 按照从低位到高位的顺序排序
	for i := 0; i < maxDigit; i++ {
		// 计算当前位的除数（10^i）
		divisor := int(math.Pow10(i))

		// 将数字分配到对应的桶中
		for _, val := range arr {
			// 取出当前位的数字
			digit := (val / divisor) % 10
			bucket[digit] = append(bucket[digit], val)
		}

		// 从桶中收集数字，更新原数组
		index := 0
		for j := 0; j < 10; j++ {
			for _, val := range bucket[j] {
				arr[index] = val
				index++
			}
			// 清空桶，准备下一轮排序
			bucket[j] = bucket[j][:0]
		}
	}
}
