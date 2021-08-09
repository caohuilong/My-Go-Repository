/*
 * @lc app=leetcode.cn id=263 lang=golang
 *
 * [263] 丑数
 */

// @lc code=start
func isUgly(n int) bool {
	if n < 1 {
		return false
	}
	for n != 1 {
		if n%2 == 0 {
			n /= 2
			continue
		} else if n%3 == 0 {
			n /= 3
			continue
		} else if n%5 == 0 {
			n /= 5
			continue
		} else {
			return false
		}

	}
	return true
}

// @lc code=end

