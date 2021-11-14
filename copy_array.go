package main

import "fmt"

func main()  {
	ans := []int{1, 2, 3}
	res := [][]int{}
	res = append(res, ans)
	ans[2] = 5
	fmt.Println(res[0])
	return
}
