package medium

type Node117 struct {
	Val   int
	Left  *Node117
	Right *Node117
	Next  *Node117
}

func connect117(root *Node117) *Node117 {
	if root == nil {
		return nil
	}
	queue := []*Node117{root}
	levelLength := 1
	for len(queue) != 0 {
		tmp := levelLength
		levelLength = 0
		for i := 0; i < tmp; i++ {
			if queue[0].Left != nil {
				queue = append(queue, queue[0].Left)
				levelLength += 1
			}
			if queue[0].Right != nil {
				queue = append(queue, queue[0].Right)
				levelLength += 1
			}
			if i != tmp-1 {
				queue[0].Next = queue[1]
			}
			queue = queue[1:]
		}
	}
	return root
}
