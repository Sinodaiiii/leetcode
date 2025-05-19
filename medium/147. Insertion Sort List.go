package medium

func insertionSortList(head *ListNode) *ListNode {
	next := head.Next
	tail := head
	for next != nil {
		if tail.Next.Val <= head.Val {
			head, tail.Next, next.Next = next, next.Next, head
			next = tail.Next
			continue
		}
		curr := head
		for curr != tail {
			if curr.Next.Val > tail.Next.Val {
				curr.Next, tail.Next, next.Next = next, next.Next, curr.Next
				break
			}
			curr = curr.Next
		}
		if tail == curr {
			tail = tail.Next
		}
		next = tail.Next
	}
	return head
}
