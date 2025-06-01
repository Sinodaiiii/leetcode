package medium

func partition(head *ListNode, x int) *ListNode {
	if head == nil {
		return nil
	}
	dummy := &ListNode{0, head}
	l := dummy
	r := l.Next
	prev := l
	for r != nil {
		if r.Val < x {
			if r == l.Next {
				l, r, prev = r, r.Next, r
				continue
			}
			l.Next, r.Next, l, r, prev.Next = r, l.Next, r, r.Next, r.Next
		} else {
			r, prev = r.Next, r
		}
	}
	return dummy.Next
}
