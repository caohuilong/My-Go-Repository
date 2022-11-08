/*
 * @lc app=leetcode.cn id=282 lang=golang
 *
 * [282] 给表达式添加运算符
 */

// @lc code=start
func addOperators(num string, target int) []string {
	ans := []string{}
	n := len(num)
	var backtrack func([]byte, int, int, int)
	backtrack = func(expr []byte, i, res, mul int) {
		if i == n {
			if res == target {
				ans = append(ans, string(expr))
			}
			return
		}

		// 符号位先添加到expr
		signIndex := len(expr)
		if i != 0 {
			expr = append(expr, ' ')
		}

		// 截取多少位加入算式，如果num[i]==0，则只能加一位到算式
		for j, val := i, 0; j < n && (i == j || num[i] != '0'); j++ {
			val = 10*val + int(num[j]-'0')
			expr = append(expr, num[j])
			if i == 0 {
				// 第一位不用加符号位
				backtrack(expr, j+1, res+val, val)
			} else {
				expr[signIndex] = '+'
				backtrack(expr, j+1, res+val, val)

				expr[signIndex] = '-'
				backtrack(expr, j+1, res-val, -val)

				expr[signIndex] = '*'
				backtrack(expr, j+1, res-mul+mul*val, mul*val)
			}
		}
	}
	backtrack(make([]byte, 0, 2*n-1), 0, 0, 0)
	return ans
}

// @lc code=end

