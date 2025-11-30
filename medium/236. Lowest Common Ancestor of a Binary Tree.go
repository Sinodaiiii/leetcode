package medium

//type TreeNode struct {
//	Val   int
//	Left  *TreeNode
//	Right *TreeNode
//}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	routeP := genRoute(root, p)
	routeQ := genRoute(root, q)
	n := min(len(routeP), len(routeQ))
	i := 0
	for i = 0; i < n; i++ {
		if routeP[i] != routeQ[i] {
			return routeQ[i-1]
		}
	}
	if len(routeP) <= len(routeQ) {
		return routeP[i-1]
	} else {
		return routeQ[i-1]
	}
}

func genRoute(root, p *TreeNode) (route []*TreeNode) {
	route = append(route, root)
	visited := map[*TreeNode]bool{}
	for len(route) != 0 {
		cur := route[len(route)-1]
		if cur == p {
			return route
		}
		if cur.Left != nil && visited[cur.Left] != true {
			route = append(route, cur.Left)
			visited[cur.Left] = true
		} else if cur.Right != nil && visited[cur.Right] != true {
			route = append(route, cur.Right)
			visited[cur.Right] = true
		} else {
			route = route[0 : len(route)-1]
		}
	}
	return nil
}
