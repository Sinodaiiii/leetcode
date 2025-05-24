package medium

func reorderList(head *ListNode) {
	if head.Next == nil {
		return
	}
	dummy := &ListNode{0, head}
	l, r := dummy, dummy
	step := 0
	for r.Next != nil && r.Next.Next != nil {
		l, r = l.Next, r.Next.Next
		step += 1
	}
	if r.Next != nil {
		l = l.Next
	}
	l.Next, l, r = nil, head, l.Next
	var prev *ListNode = nil
	for r.Next != nil {
		next := r.Next
		r.Next = prev
		prev, r = r, next
	}
	r.Next = prev
	for i := 0; i < step; i++ {
		l.Next, r.Next, l, r = r, l.Next, l.Next, r.Next
	}
}
