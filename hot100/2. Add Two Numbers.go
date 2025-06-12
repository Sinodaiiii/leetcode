package hot100

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	ans := l1
	l1, l2 = &ListNode{0, l1}, &ListNode{0, l2}
	for l1.Next != nil && l2.Next != nil {
		l1, l2 = l1.Next, l2.Next
		l1.Val = l1.Val + l2.Val
	}
	if l2.Next != nil {
		l1.Next = l2.Next
	}
	flag := 0
	head := &ListNode{0, ans}
	for head.Next != nil {
		head = head.Next
		if flag == 1 {
			head.Val += 1
			flag = 0
		}
		if head.Val/10 != 0 {
			head.Val = head.Val % 10
			flag = 1
		}
	}
	if flag == 1 {
		head.Next = &ListNode{1, nil}
	}
	return ans
}
