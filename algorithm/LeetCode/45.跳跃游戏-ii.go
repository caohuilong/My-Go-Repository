/*
 * @lc app=leetcode.cn id=45 lang=golang
 *
 * [45] 跳跃游戏 II
 */

// @lc code=start
func jump(nums []int) int {
	step := 0
	max_far := 0
	end := 0
	for i := 0; i < len(nums)-1; i++ {
		max_far = max(max_far, i+nums[i])
		if max_far >= len(nums)-1 {
			step++
			return step
		}
		if i == end {
			end = max_far
			step++
		}
	}
	return step
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
// @lc code=end

