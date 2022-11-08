/*
 * @lc app=leetcode.cn id=70 lang=golang
 *
 * [70] 爬楼梯
 */

// @lc code=start
func climbStairs(n int) int {
	if n == 1 {
		return 1
	} else if n == 2 {
		return 2
	}

	t2, t1 := 1, 2
	var t int
	for i := 3; i <= n; i++ {
		t = t1 + t2
		t2 = t1
		t1 = t
	}
	return t
}

// @lc code=end

