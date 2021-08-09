/*
 * @lc app=leetcode.cn id=480 lang=golang
 *
 * [480] 滑动窗口中位数
 */

// @lc code=start
func medianSlidingWindow(nums []int, k int) []float64 {
	var res []float64
	subNums := make([]int, k)
	for i := 0; i <= len(nums)-k; i++ {
		copy(subNums, nums[i:i+k])
		sort.Sort(sort.IntSlice(subNums))
		if k%2 == 0 {
			res = append(res, float64((subNums[k/2-1]+subNums[k/2]))/2)
		} else {
			res = append(res, float64(subNums[k/2]))
		}
	}
	return res
}

// @lc code=end

