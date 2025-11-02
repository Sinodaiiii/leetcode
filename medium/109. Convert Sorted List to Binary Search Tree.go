package medium

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sortedListToBST(head *ListNode) *TreeNode {
	if head == nil {
		return nil
	}
	dummy := &ListNode{Next: head}
	slow, fast := dummy, dummy
	for fast.Next != nil {
		fast = fast.Next
		if fast.Next != nil {
			slow = slow.Next
			fast = fast.Next
		}
	}
	slow, slow.Next = slow.Next, nil
	ret := &TreeNode{Val: slow.Val}
	ret.Left = sortedListToBST(dummy.Next)
	ret.Right = sortedListToBST(slow.Next)
	return ret
}
