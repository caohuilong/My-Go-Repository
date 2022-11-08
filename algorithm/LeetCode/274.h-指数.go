/*
 * @lc app=leetcode.cn id=274 lang=golang
 *
 * [274] H 指数
 */

// @lc code=start
func hIndex(citations []int) int {
	counts := make([]int, len(citations)+1)
	for _, i := range citations {
		if i >= len(counts) {
			counts[len(counts)-1]++
		} else {
			counts[i]++
		}
	}
	for i := len(counts)-2; i > 0; i-- {
		counts[i] += counts[i+1]
	}
	h := 0
	max := func(a, b int) int {
		if a > b {
			return a
		} else {
			return b
		}
	}
	min := func(a, b int) int {
		if a < b {
			return a
		} else {
			return b
		}
	}
	for i := len(counts)-1; i > h; i-- {
		h = max(h, min(i, counts[i]))
	}
	return h
}
// @lc code=end

