package easy

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates260110(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	pre, curr := head, head.Next
	for curr != nil {
		if curr.Val == pre.Val {
			pre.Next = curr.Next
		} else {
			pre = curr
		}
		curr = curr.Next
	}
	return head
}
