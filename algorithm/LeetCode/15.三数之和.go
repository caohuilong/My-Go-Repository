/*
 * @lc app=leetcode.cn id=15 lang=golang
 *
 * [15] 三数之和
 */

// @lc code=start
func threeSum(nums []int) [][]int {
	res := make([][]int, 0)
	sort.Ints(nums)
	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		begin, end := i+1, len(nums)-1
		for begin < end {
			if begin > i+1 && nums[begin] == nums[begin-1] {
				begin++
				continue
			}
			if end < len(nums)-1 && nums[end] == nums[end+1] {
				end--
				continue
			}
			if nums[begin]+nums[end] == -nums[i] {
				res = append(res, []int{nums[i], nums[begin], nums[end]})
				begin++
				end--
			} else if nums[begin]+nums[end] < -nums[i] {
				begin++
			} else {
				end--
			}
		}
	}
	return res
}

// @lc code=end

