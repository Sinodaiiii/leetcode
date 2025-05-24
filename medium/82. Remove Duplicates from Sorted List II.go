package medium

func deleteDuplicates(head *ListNode) *ListNode {
	dummy := &ListNode{Next: head}
	l := dummy
	r := l.Next
	for r != nil {
		if r.Next == nil || (r.Next != nil && r.Next.Val != r.Val) {
			l, r = l.Next, r.Next
			continue
		} else if r.Next != nil && r.Next.Val == r.Val {
			for r.Next != nil && r.Next.Val == r.Val {
				r = r.Next
			}
			l.Next = r.Next
			r = l.Next
		}
	}
	return dummy.Next
}
