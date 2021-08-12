/*
 * @lc app=leetcode.cn id=198 lang=golang
 *
 * [198] 打家劫舍
 */

// @lc code=start
func rob(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	} else if len(nums) == 2 {
		return max(nums[0], nums[1])
	}

	dp0, dp1 := nums[0], max(nums[0], nums[1])
	for i := 2; i < len(nums); i++ {
		tmp := max(nums[i]+dp0, dp1)
		dp0 = dp1
		dp1 = tmp
	}
	return dp1
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

// @lc code=end

