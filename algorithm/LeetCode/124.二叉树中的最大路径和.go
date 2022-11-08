/*
 * @lc app=leetcode.cn id=124 lang=golang
 *
 * [124] 二叉树中的最大路径和
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func maxPathSum(root *TreeNode) int {
	ans := math.MinInt32
	var oneSideMax func(*TreeNode) int 
	oneSideMax = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		left := max(0, oneSideMax(node.Left))
		right := max(0, oneSideMax(node.Right))
		ans = max(ans, left + right + node.Val)
		return node.Val + max(left, right)
	}
	oneSideMax(root)
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
// @lc code=end

