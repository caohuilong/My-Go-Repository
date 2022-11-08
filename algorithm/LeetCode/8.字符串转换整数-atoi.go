/*
 * @lc app=leetcode.cn id=8 lang=golang
 *
 * [8] 字符串转换整数 (atoi)
 */

// @lc code=start
func myAtoi(s string) int {
	res, sign, idx, len := 0, 1, 0, len(s)
	for idx < len && s[idx] == ' ' {
		idx++
	}
	if idx < len {
		if s[idx] == '-' {
			sign = -1
			idx++
		} else if s[idx] == '+' {
			sign = 1
			idx++
		}
	}
	for idx < len && s[idx] >= '0' && s[idx] <= '9' {
		res = res * 10 + int(s[idx]) - '0'
		if res * sign > math.MaxInt32 {
			return math.MaxInt32
		} else if sign * res < math.MinInt32 {
			return math.MinInt32
		}
		idx++
	}
	return sign * res
}
// @lc code=end

