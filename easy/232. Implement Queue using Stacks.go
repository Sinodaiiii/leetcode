package easy

type MyQueue struct {
	stack [][]int
}

func Constructor() MyQueue {
	return MyQueue{make([][]int, 2)}
}

func (this *MyQueue) Move() {
	for i := len(this.stack[0]) - 1; i >= 0; i-- {
		this.stack[1] = append(this.stack[1], this.stack[0][i])
	}
	this.stack[0] = make([]int, 0)
}

func (this *MyQueue) Push(x int) {
	this.stack[0] = append(this.stack[0], x)
}

func (this *MyQueue) Pop() int {
	if len(this.stack[1]) == 0 {
		this.Move()
	}
	ret := this.stack[1][len(this.stack[1])-1]
	this.stack[1] = this.stack[1][:len(this.stack[1])-1]
	return ret
}

func (this *MyQueue) Peek() int {
	if len(this.stack[1]) == 0 {
		this.Move()
	}
	return this.stack[1][len(this.stack[1])-1]
}

func (this *MyQueue) Empty() bool {
	if len(this.stack[0])+len(this.stack[1]) == 0 {
		return true
	}
	return false
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
