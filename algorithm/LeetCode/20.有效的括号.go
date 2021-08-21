/*
 * @lc app=leetcode.cn id=20 lang=golang
 *
 * [20] 有效的括号
 */

// @lc code=start
func isValid(s string) bool {
	bracketMap := make(map[rune]rune)
	bracketMap['('] = ')'
	bracketMap['{'] = '}'
	bracketMap['['] = ']'

	stack := make([]rune, 0)
	for _, c := range s {
		if strings.ContainsRune("({[", c) {
			stack = append(stack, c)
		} else if len(stack) == 0 || bracketMap[stack[len(stack)-1]] != c {
			return false
		} else {
			stack = stack[:len(stack)-1]
		}
	}
	if len(stack) > 0 {
		return false
	} else {
		return true
	}
}
// @lc code=end

