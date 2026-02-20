package medium

type TreeNode437 struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, targetSum int) int {
	if root == nil {
		return 0
	}
	dp := make([]int, 1001)
	ans := check(root, targetSum, dp, 0)
	return ans
}

func check(root *TreeNode, target int, dp []int, deepth int) int {
	ans := 0
	for i := deepth; i >= 0; i-- {
		dp[i] = dp[i] + root.Val
		if dp[i] == target {
			ans++
		}
	}
	if root.Left != nil {
		ans += check(root.Left, target, dp, deepth+1)
	}
	if root.Right != nil {
		ans += check(root.Right, target, dp, deepth+1)
	}
	for i := deepth; i >= 0; i-- {
		dp[i] = dp[i] - root.Val
	}
	return ans
}

func pathSum260219(root *TreeNode, targetSum int) int {
	if root == nil {
		return 0
	}
	dict := map[int]int{0: 1}
	ans := 0
	curr := 0
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		curr += node.Val
		ans += dict[curr-targetSum]
		dict[curr] += 1
		// fmt.Println(curr, dict, ans)
		if node.Left != nil {
			dfs(node.Left)
		}
		if node.Right != nil {
			dfs(node.Right)
		}
		dict[curr] -= 1
		curr -= node.Val
	}
	dfs(root)
	return ans
}
