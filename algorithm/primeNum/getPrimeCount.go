package main

func getPrimeCount(n int) int {
	primeMap := make([]bool, n+1)
	count := 0
	for i := 2; i <= n; i++ {
		if primeMap[i] == false {
			count++
			for j := i; j * i <= n; j++ {
				primeMap[i * j] = true
			}
		}
	}
	return count
}

func main()  {
	n := 6
	count := getPrimeCount(n)

	print(count)
}
