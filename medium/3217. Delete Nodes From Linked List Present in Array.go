package medium

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func modifiedList(nums []int, head *ListNode) *ListNode {
	dict := map[int]bool{}
	for _, num := range nums {
		dict[num] = true
	}
	dummy := &ListNode{Next: head}
	pre, curr := dummy, head
	for curr != nil {
		if dict[curr.Val] {
			pre.Next = curr.Next
			curr = curr.Next
		} else {
			pre = curr
			curr = curr.Next
		}
	}
	return dummy.Next
}
