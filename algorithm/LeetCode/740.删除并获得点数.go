/*
 * @lc app=leetcode.cn id=740 lang=golang
 *
 * [740] 删除并获得点数
 */

// @lc code=start
func deleteAndEarn(nums []int) int {
	sum := 0
	sort.Ints(nums)
	arr := []int{nums[0]}
	for i := 1; i < len(nums); i++ {
		val := nums[i]
		if nums[i-1] == val {
			arr[len(arr)-1] += val
		} else if nums[i-1] == val-1 {
			arr = append(arr, val)
		} else {
			sum += rob(arr)
			arr = []int{val}
		}
	}
	sum += rob(arr)
	return sum
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

