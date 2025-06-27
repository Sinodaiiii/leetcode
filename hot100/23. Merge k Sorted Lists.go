package hot100

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeKLists(lists []*ListNode) *ListNode {
	for i := len(lists) - 1; i >= 0; i-- {
		if lists[i] == nil {
			lists = append(lists[:i], lists[i+1:]...)
		}
	}
	n := len(lists)
	if n == 0 {
		return nil
	}
	var buildHeap func()
	var adjustHeap func(i int)

	buildHeap = func() {
		for i := len(lists) / 2; i >= 0; i-- {
			adjustHeap(i)
		}
	}
	adjustHeap = func(i int) {
		tmp := i
		if i*2 < len(lists) && lists[tmp].Val > lists[i*2].Val {
			tmp = i * 2
		}
		if i*2+1 < len(lists) && lists[tmp].Val > lists[i*2+1].Val {
			tmp = i*2 + 1
		}
		if tmp != i {
			lists[tmp], lists[i] = lists[i], lists[tmp]
			adjustHeap(tmp)
		}
	}

	buildHeap()
	dummy := &ListNode{}
	curr := dummy
	for len(lists) != 0 {
		adjustHeap(0)
		curr.Next = lists[0]
		curr = curr.Next
		lists[0] = lists[0].Next
		if lists[0] == nil {
			lists[0] = lists[len(lists)-1]
			lists = lists[:len(lists)-1]
		}
	}
	return dummy.Next
}
