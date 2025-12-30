package easy

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func trainingPlan(head *ListNode, cnt int) *ListNode {
	dummy := &ListNode{Next: head}
	k := 0
	for head != nil {
		k += 1
		head = head.Next
	}
	k -= cnt
	for k >= 0 {
		dummy = dummy.Next
		k -= 1
	}
	return dummy
}
