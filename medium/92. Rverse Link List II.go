package medium

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	curr := &ListNode{-1, head}
	ans := curr
	for i := 0; i < left-1; i++ {
		curr = curr.Next
	}
	first := curr
	pre := curr.Next
	if pre == nil {
		return head
	}
	post := pre.Next
	for i := 0; i < right-left && post != nil; i++ {
		post.Next, pre, post = pre, post, post.Next
	}
	first.Next.Next, first.Next = post, pre
	return ans.Next
}
