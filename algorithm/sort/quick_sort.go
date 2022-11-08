package main

import "fmt"

func quick_sort(nums []int, p, r int) {
	if p < r {
		q := partition(nums, p, r)
		quick_sort(nums, p, q-1)
		quick_sort(nums, q+1, r)
	}
}

func partition(nums []int, p, r int) int {
	x := nums[p]
	i, j := p+1, r
	for {
		for i <= j && nums[i] <= x {
			i++
		}
		for j >= i && nums[j] > x {
			j--
		}
		if i <= j {
			nums[i], nums[j] = nums[j], nums[i]
			i++
			j--
		} else {
			nums[p], nums[j] = nums[j], nums[p]
			break
		}
	}
	return j
}

func main() {
	nums := []int{10, 1, 2, 7, 9, 3, 4, 5, 6, 8}
	quick_sort(nums, 0, len(nums)-1)
	for _, num := range nums {
		fmt.Print(num, " ")
	}
	fmt.Print("\n")
	return
}
