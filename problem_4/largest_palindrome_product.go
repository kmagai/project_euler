package main

import "fmt"

// A palindromic number reads the same both ways. The largest palindrome made from the product of two 2-digit numbers is 9009 = 91 × 99.
// Find the largest palindrome made from the product of two 3-digit numbers.

func main() {
	for i := 999; i > 0; i-- {
		for r := 999; r > 0; r-- {
			n := i * r
			fmt.Println(n)
		}
	}
}
