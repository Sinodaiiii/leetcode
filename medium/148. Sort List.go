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

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func merge251228(head1, head2 *ListNode) *ListNode {
	// fmt.Println("merge: ", head1, head2)
	dummy := &ListNode{}
	ret := dummy
	for head1 != nil && head2 != nil {
		if head1.Val <= head2.Val {
			dummy.Next = head1
			dummy = dummy.Next
			head1 = head1.Next
		} else {
			dummy.Next = head2
			dummy = dummy.Next
			head2 = head2.Next
		}
	}
	if head1 != nil {
		dummy.Next = head1
	} else {
		dummy.Next = head2
	}
	return ret.Next
}

func sortList251228(head *ListNode) *ListNode {
	// fmt.Println("sort: ", head)
	if head == nil {
		return nil
	}
	if head.Next == nil {
		return head
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
	nextHead := slow.Next
	slow.Next = nil
	return merge251228(sortList(dummy.Next), sortList(nextHead))
}
