package medium

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	curr := dummy
	pre := 0
	for l1 != nil || l2 != nil {
		val1, val2 := 0, 0
		if l1 != nil {
			val1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			val2 = l2.Val
			l2 = l2.Next
		}
		curr.Next = &ListNode{Val: (val1 + val2 + pre) % 10}
		curr = curr.Next
		pre = (val1 + val2 + pre) / 10
	}
	if pre != 0 {
		curr.Next = &ListNode{Val: pre}
	}
	return dummy.Next
}
