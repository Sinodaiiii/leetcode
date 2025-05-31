package medium

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeKLists(lists []*ListNode) *ListNode {
	dummy := &ListNode{}
	cur := dummy

	var buildHeap func()
	var adjustHeap func(i int)

	buildHeap = func() {
		for i := len(lists) - 1; i >= 0; i-- {
			if lists[i] == nil {
				if i < len(lists)-1 {
					lists = append(lists[0:i], lists[i+1:]...)
				} else {
					lists = lists[0 : len(lists)-1]
				}
			}
		}
		for i := len(lists) / 2; i >= 0; i-- {
			adjustHeap(i)
		}
	}

	adjustHeap = func(i int) {
		n := len(lists)
		cur := i
		if i*2 < n && lists[cur].Val > lists[i*2].Val {
			cur = i * 2
		}
		if i*2+1 < n && lists[cur].Val > lists[i*2+1].Val {
			cur = i*2 + 1
		}
		if cur != i {
			lists[cur], lists[i] = lists[i], lists[cur]
			adjustHeap(cur)
		}
	}
	buildHeap()
	if len(lists) == 0 {
		return nil
	}

	for len(lists) != 0 {
		cur.Next, cur = lists[0], lists[0]
		lists[0] = lists[0].Next
		if lists[0] == nil {
			lists[0] = lists[len(lists)-1]
			lists = lists[:len(lists)-1]
		}
		adjustHeap(0)
	}

	return dummy.Next
}
