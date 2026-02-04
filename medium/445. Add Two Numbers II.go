package medium

import "fmt"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers260204(l1 *ListNode, l2 *ListNode) *ListNode {
	stack1, stack2 := []*ListNode{}, []*ListNode{}
	curr1, curr2 := l1, l2
	for curr1 != nil {
		stack1 = append(stack1, curr1)
		curr1 = curr1.Next
	}
	for curr2 != nil {
		stack2 = append(stack2, curr2)
		curr2 = curr2.Next
	}
	if len(stack2) > len(stack1) {
		stack1, stack2 = stack2, stack1
	}
	n1, n2 := len(stack1)-1, len(stack2)-1
	flag := 0
	for n2 >= 0 {
		stack1[n1].Val += stack2[n2].Val + flag
		if stack1[n1].Val > 9 {
			stack1[n1].Val -= 10
			flag = 1
		} else {
			flag = 0
		}
		n1 -= 1
		n2 -= 1
	}
	fmt.Println(n1)
	for flag > 0 && n1 >= 0 {
		stack1[n1].Val += flag
		if stack1[n1].Val > 9 {
			stack1[n1].Val -= 10
			flag = 1
		} else {
			flag = 0
		}
		n1 -= 1
	}
	if flag == 1 {
		return &ListNode{Val: 1, Next: stack1[0]}
	}
	return stack1[0]
}
