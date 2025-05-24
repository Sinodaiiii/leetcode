package medium

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}
	dict := map[*Node]*Node{}
	oldL := head
	newL := &Node{head.Val, nil, nil}
	dict[oldL] = newL
	ret := newL

	for newL != nil {
		if oldL.Next != nil {
			if nextNode, exist := dict[oldL.Next]; exist {
				newL.Next = nextNode
			} else {
				buildNode := &Node{oldL.Next.Val, nil, nil}
				dict[oldL.Next] = buildNode
				newL.Next = buildNode
			}
		}

		if oldL.Random != nil {
			if randNode, exist := dict[oldL.Random]; exist {
				newL.Random = randNode
			} else {
				buildNode := &Node{oldL.Random.Val, nil, nil}
				dict[oldL.Random] = buildNode
				newL.Random = buildNode
			}
		}

		oldL, newL = oldL.Next, newL.Next
	}

	return ret
}
