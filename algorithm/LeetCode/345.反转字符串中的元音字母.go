/*
 * @lc app=leetcode.cn id=345 lang=golang
 *
 * [345] 反转字符串中的元音字母
 */

// @lc code=start
func reverseVowels(s string) string {
	data := []byte(s)
	begin, end := 0, len(s)-1
	for begin < end {
		for begin < end && !strings.Contains("aeiouAEIOU", string(data[begin])) {
			begin++
		}
		for begin < end && !strings.Contains("aeiouAEIOU", string(data[end])) {
			end--
		}

		if begin < end {
			c := data[begin]
			data[begin] = data[end]
			data[end] = c
			begin++
			end--
		}
	}
	return string(data)
}

// @lc code=end

