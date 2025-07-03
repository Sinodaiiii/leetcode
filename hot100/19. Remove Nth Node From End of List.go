package hot100

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	list := []*ListNode{}
	for curr := head; curr != nil; curr = curr.Next {
		list = append(list, curr)
	}
	if n == len(list) {
		return head.Next
	} else if n == 1 {
		list[len(list)-2].Next = nil
	} else {
		list[len(list)-n-1].Next = list[len(list)-n+1]
	}
	return head
}
