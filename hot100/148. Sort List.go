package hot100

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func sortList(head *ListNode) *ListNode {
	merge := func(l1, l2 *ListNode) (first, last *ListNode) {
		newDummy := &ListNode{}
		curr := newDummy
		for l1 != nil && l2 != nil {
			if l1.Val <= l2.Val {
				curr.Next = l1
				l1 = l1.Next
			} else {
				curr.Next = l2
				l2 = l2.Next
			}
			curr = curr.Next
		}
		if l1 == nil {
			curr.Next = l2
		} else {
			curr.Next = l1
		}
		for curr.Next != nil {
			curr = curr.Next
		}
		return newDummy.Next, curr
	}
	dummy := &ListNode{0, head}
	tmp := dummy
	n := 0
	for tmp != nil {
		tmp = tmp.Next
		n += 1
	}
	for i := 1; i <= n*2; i *= 2 {
		curr := dummy
		begin := dummy
		for curr != nil {
			l1 := curr.Next
			for j := 0; j < i && curr.Next != nil; j++ {
				curr = curr.Next
			}
			l2 := curr.Next
			l1End := curr
			for j := 0; j < i && curr.Next != nil; j++ {
				curr = curr.Next
			}
			l1End.Next = nil
			l2Next := curr.Next
			curr.Next = nil
			first, last := merge(l1, l2)
			if first == nil {
				break
			} else {
				begin.Next = first
				begin = last
				begin.Next = l2Next
				curr = last
			}
		}
	}
	return dummy.Next
}
