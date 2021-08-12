/*
 * @lc app=leetcode.cn id=746 lang=golang
 *
 * [746] 使用最小花费爬楼梯
 */

// @lc code=start
func minCostClimbingStairs(cost []int) int {
	if len(cost) == 0 {
		return 0
	} else if len(cost) == 1 {
		return cost[0]
	} else if len(cost) == 2 {
		return min(cost[0], cost[1])
	}

	l := len(cost)
	fmt.print("%d\n", l)
	return min(minCostClimbingStairs(cost[:l-1]), cost[l-1]+  minCostClimbingStairs(cost[:l-2]))
}

func min(a, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}
// @lc code=end

