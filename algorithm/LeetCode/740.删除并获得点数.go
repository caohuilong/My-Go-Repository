/*
 * @lc app=leetcode.cn id=740 lang=golang
 *
 * [740] 删除并获得点数
 */

// @lc code=start
func deleteAndEarn(nums []int) int {
	maxNum := 0 
	for _, num := range nums {
		if num > maxNum {
			maxNum = num
		}
	}

	countArray := make([]int, maxNum+1)
	for _, num := range nums {
		countArray[num] += num
	}

	return rob(countArray)
}

func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	} else if len(nums) == 1 {
		return nums[0]
	}

	dp0, dp1 := 0, nums[0]
	for i := 1; i < len(nums); i++ {
		tmp := dp1
		dp1 = max(dp0+nums[i], dp1)
		dp0 = tmp
	}
	return dp1
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
// @lc code=end

