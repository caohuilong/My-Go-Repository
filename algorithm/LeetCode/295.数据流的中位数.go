/*
 * @lc app=leetcode.cn id=295 lang=golang
 *
 * [295] 数据流的中位数
 */

// @lc code=start
type MedianFinder struct {
	Nums []int
}

/** initialize your data structure here. */
func Constructor() MedianFinder {
	return MedianFinder{
		Nums: []int{},
	}
}

func (this *MedianFinder) AddNum(num int) {
	pos := this.findPosition(0, len(this.Nums)-1, num)
	this.Nums = append(this.Nums, 0)
	copy(this.Nums[0:pos], this.Nums[0:pos])
	copy(this.Nums[pos+1:], this.Nums[pos:])
	this.Nums[pos] = num
}

func (this *MedianFinder) findPosition(begin, end, num int) (pos int) {
	for begin <= end {
		mid := (begin + end) / 2
		if this.Nums[mid] < num {
			begin = mid + 1
		} else {
			end = mid - 1
		}
	}
	return begin
}

func (this *MedianFinder) FindMedian() float64 {
	length := len(this.Nums)
	if length%2 == 0 {
		return float64(this.Nums[length/2-1]+this.Nums[length/2]) / 2
	} else {
		return float64(this.Nums[length/2])
	}
}

/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */
// @lc code=end

