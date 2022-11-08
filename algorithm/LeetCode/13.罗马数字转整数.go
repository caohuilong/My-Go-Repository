/*
 * @lc app=leetcode.cn id=13 lang=golang
 *
 * [13] 罗马数字转整数
 */

// @lc code=start
func romanToInt(s string) int {
	numMap := make(map[byte]int)
	numMap['I'] = 1
	numMap['V'] = 5
	numMap['X'] = 10
	numMap['L'] = 50
	numMap['C'] = 100
	numMap['D'] = 500
	numMap['M'] = 1000
	ans := 0
	for i := 0; i < len(s); {
		if i < len(s)-1 && numMap[s[i]] < numMap[s[i+1]] {
			ans += numMap[s[i+1]] - numMap[s[i]]
			i = i + 2
		} else {
			ans += numMap[s[i]]
			i++
		}
	}
	return ans
}

// @lc code=end

