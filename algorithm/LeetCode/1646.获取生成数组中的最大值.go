/*
 * @lc app=leetcode.cn id=1646 lang=golang
 *
 * [1646] 获取生成数组中的最大值
 */

// @lc code=start
func getMaximumGenerated(n int) int {
	if n < 2 {
		return n
	}
	nums := make([]int, n+1)
	nums[0] = 0
	nums[1] = 1
	max := nums[1]
	for i := 2; i <= n; i++ {
		nums[i] = nums[i/2] + (i%2)*nums[i/2+1]
		if nums[i] > max {
			max = nums[i]
		}
	}
	return max
}

// @lc code=end

