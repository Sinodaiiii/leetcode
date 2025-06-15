package hot100

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	curr := &ListNode{}
	ret := curr
	for list1 != nil && list2 != nil {
		if list1.Val > list2.Val {
			curr.Next = list2
			curr, list2 = list2, list2.Next
		} else {
			curr.Next = list1
			curr, list1 = list1, list1.Next
		}
	}
	if list1 != nil {
		curr.Next = list1
	} else {
		curr.Next = list2
	}
	return ret.Next
}
