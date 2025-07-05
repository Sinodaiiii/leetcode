package hot100

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func hasCycle(head *ListNode) bool {
	dummy := &ListNode{0, head}
	slow, fast := dummy, head
	for slow != nil && fast != nil {
		slow = slow.Next
		fast = fast.Next
		if fast == nil {
			return false
		} else if fast == slow || fast.Next == slow {
			return true
		}
		fast = fast.Next
	}
	return false
}
