package main

import "fmt"

// The prime factors of 13195 are 5, 7, 13 and 29.
// What is the largest prime factor of the number 600851475143 ?

func main() {
	n := 600851475143
	fmt.Println(max(primeFactor(n)))
}

func primeFactor(n int) []int {
	m := []int{}
	for i := 2; i <= n; i++ {
		if n%i == 0 {
			n = n / i
			m = append(m, i)
		}
	}
	return m
}

func max(arr []int) int {
	max := arr[0]
	for _, i := range arr {
		if i > max {
			max = i
		}
	}
	return max
}
