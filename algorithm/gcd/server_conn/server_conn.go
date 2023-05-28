package main

import (
	"fmt"
)

type Server struct {
	Name    string
	Weight  int
	CurConn int
}

// 计算服务器应该处理的连接数
func calcConn(servers []Server, totalConn int) []int {
	sum := 0
	for _, s := range servers {
		sum += s.Weight
	}

	nums := make([]int, len(servers))
	for i, s := range servers {
		nums[i] = s.Weight * totalConn / sum
	}

	return nums
}

// 重新分配连接数
func rebalanceConn(servers []Server, nums []int) {
	servers = sortServers(servers, nums)

	maxPos := len(servers) - 1
	curPos := 0

	for curPos <= maxPos {
		s := &servers[curPos]
		if s.CurConn > nums[curPos] {
			// 需要将连接数分配给其他服务器
			diff := s.CurConn - nums[curPos]

			for i := maxPos; i > curPos && diff > 0; i-- {
				s2 := &servers[i]
				if s2.CurConn < nums[i] {
					if nums[i]-s2.CurConn >= diff {
						s2.CurConn += diff
						s.CurConn -= diff
						diff = 0
					} else {
						diff -= nums[i] - s2.CurConn
						s.CurConn -= nums[i] - s2.CurConn
						s2.CurConn = nums[i]
					}
				}
			}
		}

		curPos++
	}
}

// 根据权重和处理相应连接数排序服务器
func sortServers(servers []Server, nums []int) []Server {
	newServers := make([]Server, len(servers))
	copy(newServers, servers)

	for i := 0; i < len(newServers); i++ {
		for j := i + 1; j < len(newServers); j++ {
			if nums[j-1]*newServers[j].Weight > nums[j]*newServers[j-1].Weight {
				newServers[j-1], newServers[j] = newServers[j], newServers[j-1]
				nums[j-1], nums[j] = nums[j], nums[j-1]
			}
		}
	}

	return newServers
}

func main() {
	servers := []Server{
		{Name: "Server A", Weight: 2, CurConn: 25},
		{Name: "Server B", Weight: 3, CurConn: 35},
		{Name: "Server C", Weight: 5, CurConn: 40},
	}

	fmt.Println("Original servers:", servers)

	// 计算服务器应该处理的连接数
	nums := calcConn(servers, 100)

	// 重新分配连接数
	rebalanceConn(servers, nums)

	fmt.Println("Rebalanced servers:", servers)
}