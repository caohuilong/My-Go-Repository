/*
 * @lc app=leetcode.cn id=541 lang=golang
 *
 * [541] 反转字符串 II
 */

// @lc code=start
func reverseStr(s string, k int) string {
	data := []byte(s)
	for i := 0; i < len(s); i += 2 * k {
		n := min(i+k, len(s))
		for begin, end := i, n-1; begin < end; begin, end = begin+1, end-1 {
			data[begin], data[end] = data[end], data[begin]
		}
	}
	return string(data)
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

// @lc code=end

