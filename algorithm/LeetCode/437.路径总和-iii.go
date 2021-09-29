/*
 * @lc app=leetcode.cn id=437 lang=golang
 *
 * [437] 路径总和 III
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

// 前缀和
func pathSum(root *TreeNode, targetSum int) int {
	ans := 0
	prefix := map[int]int{0: 1}
	var dfs func(*TreeNode, int)
	dfs = func(node *TreeNode, curr int) {
		if node == nil {
			return
		}
		curr += node.Val
		ans += prefix[curr-targetSum]
		prefix[curr]++
		dfs(node.Left, curr)
		dfs(node.Right, curr)
		prefix[curr]--
	}
	dfs(root, 0)
	return ans
}

/** 深度优先搜索
func rootSum(root *TreeNode, targetSum int) int {
	if root == nil {
		return 0
	}
	res := 0
	if root.Val == targetSum {
		res += 1
	}
	res += rootSum(root.Left, targetSum-root.Val)
	res += rootSum(root.Right, targetSum-root.Val)
	return res
}

func pathSum(root *TreeNode, targetSum int) int {
	if root == nil {
		return 0
	}
	res := rootSum(root, targetSum)
	res += pathSum(root.Left, targetSum)
	res += pathSum(root.Right, targetSum)
	return res
}
*/

// @lc code=end

