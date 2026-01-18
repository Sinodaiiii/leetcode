package main

import (
	"leetcode/goroutine"
)

//type heapSlice []int
//
//func (h heapSlice) Len() int {
//	return len(h)
//}
//
//func (h heapSlice) Less(i, j int) bool {
//	return h[i] < h[j]
//}
//
//func (h heapSlice) Swap(i, j int) {
//	h[i], h[j] = h[j], h[i]
//}
//
//func (h *heapSlice) Pop() interface{} {
//	num := *h
//	ret := num[len(num)-1]
//	*h = num[0 : len(*h)-1]
//	return ret
//}
//
//func (h *heapSlice) Push(i interface{}) {
//	*h = append(*h, i.(int))
//}

func main() {
	goroutine.Print3Thread()
	//t := []int{5, 4, 10, 20, 100}
	//ds.QuickSort(t, 0, len(t)-1)
	//ds.HeapSort(t)
	//fmt.Println(t)
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

}
