/*
 * @lc app=leetcode.cn id=1175 lang=golang
 *
 * [1175] 质数排列
 */

// @lc code=start
func numPrimeArrangements(n int) int {
	if n == 1 {
		return 1
	}
	count := getPrimeCount(n)
	return (jieChen(count) * jieChen(n-count)) % 1000000007
}

func getPrimeCount(n int) int {
	primeMap := make([]bool, n+1)
	count := 0
	for i := 2; i <= n; i++ {
		if primeMap[i] == false {
			count++
			for j := 2; i*j <= n; j++ {
				primeMap[i*j] = true
			}
		}
	}
	return count
}

func jieChen(n int) int {
	ans := 1
	for i := n; i > 1; i-- {
		ans = (ans * i) % (1e9 + 7)
	}
	return ans
}

// @lc code=end

