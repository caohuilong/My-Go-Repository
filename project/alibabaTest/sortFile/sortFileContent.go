package main

import (
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

func main() {
	bytes, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}
	strsList := strings.Split(string(bytes), "\n")

	sort.Slice(strsList, func(i, j int) bool {
		nums1 := strings.Split(strsList[i], ".")
		nums2 := strings.Split(strsList[j], ".")
		for k := 0; k < len(nums1) && k < len(nums2); k = k + 1 {
			n1, _ := strconv.Atoi(nums1[k])
			n2, _ := strconv.Atoi(nums2[k])
			if n1 < n2 {
				return true
			} else if n1 > n2 {
				return false
			}
		}
		if len(nums1) < len(nums2) {
			return true
		}
		return false
	})
	err = ioutil.WriteFile("./output.txt", []byte(strings.Join(strsList, "\n")), 0644)
	if err != nil {
		panic(err)
	}
	return
}
