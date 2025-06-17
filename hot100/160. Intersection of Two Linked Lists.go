package hot100

// space complexity can be optimized to o(1)
// traverse m+n list, return the first same ListNode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	dict := map[*ListNode]bool{}
	for headA != nil {
		dict[headA] = true
		headA = headA.Next
	}
	for headB != nil {
		if dict[headB] == true {
			return headB
		}
		headB = headB.Next
	}
	return nil
}
