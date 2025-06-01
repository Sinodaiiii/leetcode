package problems

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	if k == 1 {
		return head
	}
	dummy := &ListNode{0, head}
	count := 1
	var pre, cur *ListNode = dummy, head
	for head != nil {
		head = head.Next
		count += 1
		if count == k && head != nil {
			head = head.Next
			first := pre
			for cur != head {
				pre, cur, cur.Next = cur, cur.Next, pre
			}
			first.Next, first.Next.Next, pre = pre, head, first.Next
			count = 1
		}
	}
	return dummy.Next
}
