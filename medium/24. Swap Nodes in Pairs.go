package medium

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapPairs260109(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	dummy := &ListNode{Next: head}
	pre, curr, next := dummy, head, head.Next
	for curr != nil && next != nil {
		pre.Next, curr.Next, next.Next = next, next.Next, curr
		if curr.Next == nil {
			break
		}
		pre, curr, next = curr, curr.Next, curr.Next.Next
	}
	return dummy.Next
}
