package main

import "fmt"

// A palindromic number reads the same both ways. The largest palindrome made from the product of two 2-digit numbers is 9009 = 91 × 99.
// Find the largest palindrome made from the product of two 3-digit numbers.

func main() {
	max := 0
	for i := 999; i > 0; i-- {
		for r := 999; r > 0; r-- {
			if n := i * r; n == reverse(n) && n > max {
				max = n
			}
		}
	}
	fmt.Println(max)
}

func reverse(n int) int {
	acm := 0
	for ; n > 1; n = n / 10 {
		acm = acm*10 + n%10
	}
	return acm
}
