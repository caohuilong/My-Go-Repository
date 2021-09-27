/*
 * @lc app=leetcode.cn id=639 lang=golang
 *
 * [639] 解码方法 II
 */

// @lc code=start
func check1digit(c byte) int {
	if c == '*' {
		return 9
	} else if c == '0' {
		return 0
	}
	return 1
}

func check2digits(c1, c2 byte) int {
	if c1 == '*' && c2 == '*' {
		return 15
	}
	if c1 == '*' {
		if c2 <= '6' {
			return 2
		} else {
			return 1
		}
	}
	if c1 == '1' {
		if c2 == '*' {
			return 9
		} else {
			return 1
		}
	}
	if c1 == '2' {
		if c2 == '*' {
			return 6
		} else if c2 <= '6' {
			return 1
		}
	}
	return 0
}

func numDecodings(s string) int {
	dp1, dp2, dp := 0, 1, 0
	const mod int = 1e9 + 7
	for i := range s {
		dp = (dp2 * check1digit(s[i])) % mod
		if i >= 1 {
			dp = (dp + dp1*check2digits(s[i-1], s[i])) % mod
		}
		dp1, dp2 = dp2, dp
	}
	return dp
}

// @lc code=end

