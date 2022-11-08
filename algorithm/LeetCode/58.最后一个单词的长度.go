/*
 * @lc app=leetcode.cn id=58 lang=golang
 *
 * [58] 最后一个单词的长度
 */

// @lc code=start
func lengthOfLastWord(s string) int {
	str := []byte(s)
	count := 0
	for i := len(str) - 1; i >= 0; i-- {
		if str[i] == ' ' && count > 0 {
			break
		} else if str[i] != ' ' {
			count++
		}
	}
	return count
}

// @lc code=end

