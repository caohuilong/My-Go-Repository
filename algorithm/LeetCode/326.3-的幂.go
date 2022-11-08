/*
 * @lc app=leetcode.cn id=326 lang=golang
 *
 * [326] 3的幂
 */

// @lc code=start
func isPowerOfThree(n int) bool {
	for n > 0 && n%3 == 0 {
		n /= 3
	}
	return n == 1
}

// @lc code=end

