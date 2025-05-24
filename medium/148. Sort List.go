package medium

import "sort"

func sortList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	l := make([]*ListNode, 0)
	for head != nil {
		l = append(l, head)
		head = head.Next
	}
	sort.Slice(l, func(i, j int) bool { return l[i].Val <= l[j].Val })
	l[len(l)-1].Next = nil
	for i := len(l) - 2; i >= 0; i-- {
		l[i].Next = l[i+1]
	}
	return l[0]
}
