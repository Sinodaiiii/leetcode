package hot100

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
	var priv *ListNode = nil
	for head != nil {
		priv, head, head.Next = head, head.Next, priv
	}
	return priv
}
