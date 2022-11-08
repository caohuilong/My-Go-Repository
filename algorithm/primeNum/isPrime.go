package main

import "math"

func isPrime(num int) bool {
	if num <= 3 {
		return num > 1
	}
	if num%6 != 1 && num%6 != 5 {
		return false
	}

	sqrt := int(math.Sqrt(float64(num)))
	for i := 5; i <= sqrt; i+=6 {
		if num%i == 0 || num%i == 2 {
			return false
		}
	}
	return true
}

func main()  {
	print(isPrime(53))
	return
}