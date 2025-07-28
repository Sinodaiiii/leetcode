package medium

type DoublyNode struct {
	val        int
	prev, next *DoublyNode
}

type MyLinkedList struct {
	length     int
	head, tail *DoublyNode
}

func Constructor707() MyLinkedList {
	head := &DoublyNode{}
	tail := &DoublyNode{}
	head.next = tail
	tail.prev = head
	return MyLinkedList{0, head, tail}
}

func (this *MyLinkedList) Get(index int) int {
	if index >= this.length {
		return -1
	}
	curr := this.head.next
	for index > 0 {
		index -= 1
		curr = curr.next
	}
	return curr.val
}

func (this *MyLinkedList) AddAtHead(val int) {
	this.length += 1
	tmp := &DoublyNode{val, this.head, this.head.next}
	this.head.next, this.head.next.prev = tmp, tmp
}

func (this *MyLinkedList) AddAtTail(val int) {
	this.length += 1
	tmp := &DoublyNode{val, this.tail.prev, this.tail}
	this.tail.prev, this.tail.prev.next = tmp, tmp
}

func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if index > this.length {
		return
	}
	curr := this.head
	for index > 0 {
		index -= 1
		curr = curr.next
	}
	tmp := &DoublyNode{val, curr, curr.next}
	curr.next, curr.next.prev = tmp, tmp
	this.length += 1
}

func (this *MyLinkedList) DeleteAtIndex(index int) {
	if index < 0 || index >= this.length {
		return
	}
	curr := this.head.next
	for index > 0 {
		index -= 1
		curr = curr.next
	}
	curr.prev.next, curr.next.prev = curr.next, curr.prev
	this.length -= 1
}

/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */
