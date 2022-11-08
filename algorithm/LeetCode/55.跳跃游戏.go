/*
 * @lc app=leetcode.cn id=55 lang=golang
 *
 * [55] 跳跃游戏
 */

// @lc code=start
func canJump(nums []int) bool {
	remote := 0
	for i := 0; i < len(nums); i++ {
		if i > remote {
			return false
		}
		tmp := i + nums[i]
		if tmp >= len(nums)-1 {
			return true
		} else if tmp > remote {
			remote = tmp
		}
	}
	return false
}
// @lc code=end

