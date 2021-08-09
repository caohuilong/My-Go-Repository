/*
 * @lc app=leetcode.cn id=3 lang=golang
 *
 * [3] 无重复字符的最长子串
 */

// @lc code=start
func lengthOfLongestSubstring(s string) int {
	data := []byte(s)
	charMap := make(map[byte]bool)
	rk := -1
	res := 0
	i := 0
	for i < len(data) {
		for rk+1 < len(data) && !charMap[data[rk+1]] {
			charMap[data[rk+1]] = true
			rk++
		}
		res = max(res, len(charMap))
		j := i
		for !charMap[data[j]] {
			delete(charMap, data[j])
			j++
		}
		i = j
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

// @lc code=end

