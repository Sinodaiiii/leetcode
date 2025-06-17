package hot100

// space complexity can be optimized to o(1)
// first find the middle point (by slow-fast pointer)
// reverse the sencond half
// compare

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func isPalindrome(head *ListNode) bool {
	stack := make([]int, 0)
	for head != nil {
		stack = append(stack, head.Val)
		head = head.Next
	}
	n := len(stack)
	for i := 0; i <= n/2; i++ {
		if stack[i] != stack[n-1-i] {
			return false
		}
	}
	return true
}
