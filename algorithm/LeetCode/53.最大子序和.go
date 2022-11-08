/*
 * @lc app=leetcode.cn id=53 lang=golang
 *
 * [53] 最大子序和
 */

// @lc code=start
func maxSubArray(nums []int) int {
	pre := nums[0]
	ans := nums[0]
	for i := 1; i < len(nums); i++ {
		pre = max(pre+nums[i], nums[i])
		ans = max(pre, ans)
	}
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

