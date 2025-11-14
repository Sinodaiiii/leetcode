package medium

type Node133 struct {
	Val       int
	Neighbors []*Node133
}

func cloneGraph(node *Node133) *Node133 {
	if node == nil {
		return nil
	}
	queue := []*Node133{node}
	dict := map[int]*Node133{}
	visited := map[int]bool{node.Val: true}
	next := map[int][]int{}
	for len(queue) != 0 {
		curr := queue[0]
		dict[curr.Val] = &Node133{curr.Val, nil}
		for _, nei := range curr.Neighbors {
			next[curr.Val] = append(next[curr.Val], nei.Val)
			if !visited[nei.Val] {
				visited[nei.Val] = true
				queue = append(queue, nei)
			}
		}
		queue = queue[1:]
	}
	for k, v := range next {
		curr := dict[k]
		for _, nei := range v {
			curr.Neighbors = append(curr.Neighbors, dict[nei])
		}
	}
	return dict[node.Val]
}
