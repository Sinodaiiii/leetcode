package medium

type ListNode struct {
	Val  int
	Next *ListNode
}

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil {
		return head
	}
	l := 0
	var first, pre *ListNode
	first = head
	for head != nil {
		l++
		pre = head
		head = head.Next
	}
	pre.Next = first
	k = l - k%l
	head = pre
	for k > 0 {
		k--
		head = head.Next
	}
	ans := head.Next
	head.Next = nil
	return ans
}
