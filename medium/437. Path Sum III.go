package medium

type TreeNode struct {
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
