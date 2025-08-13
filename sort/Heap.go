package sort

import "fmt"

// MaxHeap 最大堆结构体
type MaxHeap struct {
	data []int
}

// NewMaxHeap 创建新的最大堆
func NewMaxHeap() *MaxHeap {
	return &MaxHeap{data: []int{}}
}

// Insert 向堆中插入元素
func (h *MaxHeap) Insert(val int) {
	h.data = append(h.data, val)
	h.heapifyUp(len(h.data) - 1)
}

// heapifyUp 向上调整堆
func (h *MaxHeap) heapifyUp(index int) {
	for index > 0 {
		parent := (index - 1) / 2
		if h.data[index] <= h.data[parent] {
			break
		}
		h.data[index], h.data[parent] = h.data[parent], h.data[index]
		index = parent
	}
}

// BuildHeap 从切片构建最大堆
func (h *MaxHeap) BuildHeap(arr []int) {
	h.data = arr
	n := len(h.data)
	for i := n/2 - 1; i >= 0; i-- {
		h.adjustHeap(i, n)
	}
}

// adjustHeap 向下调整堆（维护最大堆性质）
func (h *MaxHeap) adjustHeap(index, heapSize int) {
	largest := index
	left := 2*index + 1
	right := 2*index + 2

	if left < heapSize && h.data[left] > h.data[largest] {
		largest = left
	}

	if right < heapSize && h.data[right] > h.data[largest] {
		largest = right
	}

	if largest != index {
		h.data[index], h.data[largest] = h.data[largest], h.data[index]
		h.adjustHeap(largest, heapSize)
	}
}

// GetMax 获取堆顶元素（最大值）
func (h *MaxHeap) GetMax() int {
	if len(h.data) == 0 {
		return -1 // 空堆处理
	}
	return h.data[0]
}

// Pop 移除并返回堆顶元素（最大值）
func (h *MaxHeap) Pop() int {
	n := len(h.data)
	if n == 0 {
		return -1 // 空堆处理
	}

	max := h.data[0]
	// 将最后一个元素移到堆顶
	h.data[0] = h.data[n-1]
	h.data = h.data[:n-1]        // 缩小数组
	h.adjustHeap(0, len(h.data)) // 向下调整堆
	return max
}

func main() {
	heap := NewMaxHeap()
	heap.Insert(3)
	heap.Insert(5)
	heap.Insert(1)
	heap.Insert(10)

	fmt.Println("堆顶元素:", heap.GetMax()) // 输出: 10
	fmt.Println("弹出元素:", heap.Pop())    // 输出: 10
	fmt.Println("剩余堆顶:", heap.GetMax()) // 输出: 5
}

// 迭代的push后向上、向下调整
// 大顶堆向上调整（用于添加元素到堆尾后）
func maxHeapUp(maxHeap []int, index int) {
	for index > 0 {
		parent := (index - 1) / 2
		if maxHeap[parent] >= maxHeap[index] {
			break // 父节点更大，无需调整
		}
		// 交换父节点和当前节点
		maxHeap[parent], maxHeap[index] = maxHeap[index], maxHeap[parent]
		index = parent
	}
}

// 大顶堆向下调整（用于堆顶元素变化后）
func maxHeapDown(maxHeap []int, index int) {
	length := len(maxHeap)
	for {
		left := 2*index + 1
		right := 2*index + 2
		largest := index

		// 找到左、右子节点中更大的那个
		if left < length && maxHeap[left] > maxHeap[largest] {
			largest = left
		}
		if right < length && maxHeap[right] > maxHeap[largest] {
			largest = right
		}

		if largest == index {
			break // 无需调整
		}
		// 交换当前节点和最大子节点
		maxHeap[index], maxHeap[largest] = maxHeap[largest], maxHeap[index]
		index = largest
	}
}
