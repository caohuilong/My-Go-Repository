/*
 * @lc app=leetcode.cn id=202 lang=golang
 *
 * [202] 快乐数
 */

// @lc code=start
func isHappy(n int) bool {
	numMap := make(map[int]bool, 0)
	for n > 0 {
		if n == 1 {
			return true
		}
		if _, isExist := numMap[n]; isExist {
			return false
		} else {
			numMap[n] = true
		}
		s := 0
		for n > 0 {
			k := n % 10
			s += k * k
			n /= 10
		}
		n = s
	}
	return false
}

// @lc code=end

