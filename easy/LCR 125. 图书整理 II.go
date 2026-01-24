package easy

type CQueue struct {
	inStack  []int
	outStack []int
}

func Constructor260125() CQueue {
	return CQueue{make([]int, 0, 10000), make([]int, 0, 10000)}
}

func (this *CQueue) AppendTail(value int) {
	this.inStack = append(this.inStack, value)
}

func (this *CQueue) DeleteHead() int {
	if len(this.outStack) == 0 {
		for i := len(this.inStack) - 1; i >= 0; i-- {
			this.outStack = append(this.outStack, this.inStack[i])
		}
		this.inStack = this.inStack[:0]
	}
	if len(this.outStack) == 0 {
		return -1
	}
	ret := this.outStack[len(this.outStack)-1]
	this.outStack = this.outStack[:len(this.outStack)-1]
	return ret
}

/**
 * Your CQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AppendTail(value);
 * param_2 := obj.DeleteHead();
 */
