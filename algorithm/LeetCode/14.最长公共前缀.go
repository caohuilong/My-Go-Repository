/*
 * @lc app=leetcode.cn id=14 lang=golang
 *
 * [14] 最长公共前缀
 */

// @lc code=start
func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	prefix := strs[0]
	for i := 1; i < len(strs); i++ {
		prefix = lcp(prefix, strs[i])
		if len(prefix) == 0 {
			return ""
		}
	}
	return prefix
}

func lcp(str1, str2 string) string {
	len := min(len(str1), len(str2))
	index := 0
	for ; index < len; index++ {
		if str1[index] != str2[index] {
			return str1[:index]
		}
	}
	return str1[:index]
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

// @lc code=end

