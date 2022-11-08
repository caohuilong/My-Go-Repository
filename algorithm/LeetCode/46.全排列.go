/*
 * @lc app=leetcode.cn id=46 lang=golang
 *
 * [46] 全排列
 */

// @lc code=start

var res [][]int
func permute(nums []int) [][]int {
	res = make([][]int, 0)
	trace := make([]int, 0)
	backTrace(nums, trace)
	return res
}

func backTrace(nums []int, trace []int)  {
	if len(trace) == len(nums) {
		res = append(res, trace)
		return
	}
	for i := 0; i < len(nums); i++ {
		if isInSlice(nums[i], trace) {
			continue
		}
		trace = append(trace, nums[i])
		backTrace(nums, trace)
		trace = trace[:len(trace)-1]
	}
}

func isInSlice(target int, trace []int) bool {
	for _, num := range trace {
		if num == target {
			return true
		}
	}
	return false
}
// @lc code=end

