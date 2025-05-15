package medium

import "math"

type MinStack struct {
	minv  int
	mini  []int
	stack []int
}

func Constructor() MinStack {
	return MinStack{math.MaxInt32, make([]int, 0), make([]int, 0)}
}

func (this *MinStack) Push(val int) {
	this.stack = append(this.stack, val)
	if val <= this.minv {
		i := len(this.stack)
		this.minv = val
		this.mini = append(this.mini, i-1)
	}
}

func (this *MinStack) Pop() {
	i := len(this.stack) - 1
	if this.mini[len(this.mini)-1] == i {
		this.mini = this.mini[:len(this.mini)-1]
		if len(this.mini) == 0 {
			this.minv = math.MaxInt32
		} else {
			minStacki := this.mini[len(this.mini)-1]
			this.minv = this.stack[minStacki]
		}
	}
	this.stack = this.stack[:i]
}

func (this *MinStack) Top() int {
	return this.stack[len(this.stack)-1]
}

func (this *MinStack) GetMin() int {
	return this.minv
}
