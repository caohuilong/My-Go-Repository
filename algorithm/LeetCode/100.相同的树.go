/*
 * @lc app=leetcode.cn id=100 lang=golang
 *
 * [100] 相同的树
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

// 深度优先
func isSameTree0(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	} else if p == nil && q != nil {
		return false
	} else if p != nil && q == nil {
		return false
	} else if p.Val != q.Val {
		return false
	}
	return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}

// 广度优先
func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	} else if p == nil || q == nil {
		return false
	}
	q1, q2 := []*TreeNode{p}, []*TreeNode{q}
	for len(q1) != 0 && len(q2) != 0 {  
		a, b := q1[0], q2[0]
		q1, q2 = q1[1:], q2[1:]
		if a.Val != b.Val {
			return false
		}
		left1, right1 := a.Left, a.Right
		left2, right2 := b.Left, b.Right
		if left1 == nil && left2 != nil || left1 != nil && left2 == nil {
			return false
		}
		if right1 == nil && right2 != nil || right1 != nil && right2 == nil {
			return false
		}
		if left1 != nil {
			q1 = append(q1, left1)
		}
		if right1 != nil {
			q1 = append(q1, right1)
		}
		if left2 != nil {
			q2 = append(q2, left2)
		}
		if right2 != nil {
			q2 = append(q2, right2)
		}
	}
	if len(q1) !=0 || len(q2) != 0 {
		return false
	}
	return true
}
// @lc code=end

