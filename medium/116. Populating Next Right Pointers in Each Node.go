package medium

type Node116 struct {
	Val   int
	Left  *Node116
	Right *Node116
	Next  *Node116
}

func connect(root *Node116) *Node116 {
	if root == nil {
		return nil
	}
	queue := []*Node116{root}
	levelLength := 1
	for len(queue) != 0 {
		for i := 0; i < levelLength; i++ {
			if queue[0].Left != nil {
				queue = append(queue, queue[0].Left)
				queue = append(queue, queue[0].Right)
			}
			if i != levelLength-1 {
				queue[0].Next = queue[1]
			}
			queue = queue[1:]
		}
		levelLength *= 2
	}
	return root
}
