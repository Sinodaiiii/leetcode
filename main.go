package main

import "fmt"

type heapSlice []int

func (h heapSlice) Len() int {
	return len(h)
}

func (h heapSlice) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h heapSlice) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *heapSlice) Pop() interface{} {
	num := *h
	ret := num[len(num)-1]
	*h = num[0 : len(*h)-1]
	return ret
}

func (h *heapSlice) Push(i interface{}) {
	*h = append(*h, i.(int))
}

func main() {
	//println("hello")
	//h := heapSlice{5, 4, 3, 2, 1}
	//heap.Init(&h)
	//for i := 0; i < len(h); i++ {
	//	println(h[i])
	//}
	//heap.Pop(&h)
	//for i := 0; i < len(h); i++ {
	//	println(h[i])
	//}
	//heap.Push(&h, 0)
	//for i := 0; i < len(h); i++ {
	//	println(h[i])
	//}
	//sort.Slice(h, func(i, j int) bool {
	//	return h[i] > h[j]
	//})
	//for i := 0; i < len(h); i++ {
	//	println(h[i])
	//}
	//
	//var m, n int
	//fmt.Scanf("%d %d", &m, &n)
	//
	//// 初始化二维切片（矩阵）
	//t := make([][]byte, m)
	//for i := range t {
	//	t[i] = make([]byte, n)
	//}
	//
	//// 读取矩阵内容
	//for i := 0; i < m; i++ {
	//	for j := 0; j < n; j++ {
	//		var c byte
	//		fmt.Scanf("%c", &c) // 读取单个字符
	//		t[i][j] = c
	//	}
	//	// 跳过行尾的换行符（如果输入是每行回车分隔）
	//	fmt.Scanf("\n")
	//}
	//
	//// 正确输出矩阵
	//for i := 0; i < m; i++ {
	//	for j := 0; j < n; j++ {
	//		fmt.Printf("%c ", t[i][j])
	//	}
	//	fmt.Println()
	//}

	t := [][]byte{}
	m, n := 0, 0
	fmt.Scanf("%d %d", &m, &n)
	t = make([][]byte, m)
	for i := 0; i < m; i++ {
		t[i] = make([]byte, 0)
		for j := 0; j < n; j++ {
			in := '0'
			fmt.Scanf("%c", &in)
			t[i] = append(t[i], byte(in))
		}
	}
	print(t)

	input1 := "[1,[2,[3,[4,5]]]]"
	input2 := "[1,[2,[3,[4,5,6]]]]"

	reverse := func(input string) string {
		inputSlice := []byte(input)
		i, j := 0, len(inputSlice)-1
		for {
			for input[i] < '0' || input[i] > '9' {
				i++
			}
			for input[j] < '0' || input[j] > '9' {
				j--
			}
			if i >= j {
				break
			}
			inputSlice[i], inputSlice[j] = inputSlice[j], inputSlice[i]
			i++
			j--
		}
		return string(inputSlice)
	}

	output1 := reverse(input1)
	output2 := reverse(input2)
	print(output1, output2)
}
