package main

import "fmt"

type Server struct {
	Name    string
	Weight  int
	CurConn int
}

// 根据权重排序服务器
func sortServers(servers []Server) []Server {
	newServers := make([]Server, len(servers))
	copy(newServers, servers)

	for i := 0; i < len(newServers); i++ {
		for j := i + 1; j < len(newServers); j++ {
			if newServers[i].Weight < newServers[j].Weight {
				newServers[i], newServers[j] = newServers[j], newServers[i]
			}
		}
	}

	return newServers
}

// 获取所有权重的最大公约数
func getGCD(weights []int) int {
	num := len(weights)
	if num == 1 {
		return weights[0]
	} else {
		a := weights[0]
		for i := 1; i < num; i++ {
			b := weights[i]
			a = gcd(a, b)
		}
		return a
	}
}

// 计算最大公约数
func gcd(a, b int) int {
	if b == 0 {
		return a
	} else {
		return gcd(b, a%b)
	}
}

// 根据权重计算每个服务器要处理的请求数量
func getNums(weights []int, gcd int) []int {
	nums := make([]int, len(weights))
	for i, w := range weights {
		nums[i] = w / gcd
	}
	return nums
}

// 分配请求给服务器
func dispatchRequest(servers []Server, nums []int) {
	maxPos := len(servers)
	curPos := 0

	for {
		if curPos == maxPos {
			curPos = 0
		}

		if nums[curPos] > 0 {
			s := &servers[curPos]
			s.CurConn++
			fmt.Printf("Request sent to server %s (weight: %d, connections: %d)\n", s.Name, s.Weight, s.CurConn)

			nums[curPos]--
		}

		curPos++
	}
}

func main() {
	servers := []Server{
		{Name: "Server A", Weight: 3},
		{Name: "Server B", Weight: 2},
		{Name: "Server C", Weight: 1},
	}

	servers = sortServers(servers)

	weights := make([]int, len(servers))
	for i, s := range servers {
		weights[i] = s.Weight
	}

	gcdNum := getGCD(weights)
	nums := getNums(weights, gcdNum)

	fmt.Println("Server sorted by weight:", servers)
	fmt.Println("Max GCD:", gcdNum)
	fmt.Println("Nums per server:", nums)

	dispatchRequest(servers, nums)
}