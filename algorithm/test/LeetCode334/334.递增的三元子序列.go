package main

import "fmt"

func main() {
	nums := []int{4, 5, 6, 1, 2}
	res := increasingTriplet(nums)
	fmt.Println(res)
}

func increasingTriplet(nums []int) bool {
	n := len(nums)
	minList, maxList := make([]int, n), make([]int, n)

	for i := 0; i < n; i++ {
		if i == 0 {
			minList[i] = nums[i]
			maxList[n-i-1] = nums[n-i-1]
		} else {
			if nums[i-1] < minList[i-1] {
				minList[i] = nums[i-1]
			} else {
				minList[i] = minList[i-1]
			}
			if nums[n-i] > maxList[n-i] {
				maxList[n-i-1] = nums[n-i]
			} else {
				maxList[n-i-1] = maxList[n-i]
			}
		}
	}

	for i := 1; i < n-1; i++ {
		if nums[i] > minList[i] && nums[i] < maxList[i] {
			return true
		}
	}
	return false
}
