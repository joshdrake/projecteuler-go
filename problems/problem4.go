// Find the largest palindrome made from the product of two 3-digit numbers.

package main

import (
	"fmt"
	"strconv"
)

func main() {
	max := 0
	for i := 999; i > 99; i-- {
		for j := 999; j > 99; j-- {
			p := i * j
			p_str := strconv.Itoa(p)
			if p_str == reversed(p_str) {
				if p > max {
					max = p
				}
			}
		}
	}

	fmt.Println(max)
}

func reversed(s string) string {
	strlen := len(s)
	a := make([]byte, strlen)
	for i := 0; i < strlen; i++ {
		a[strlen-i-1] = s[i]
	}
	return string(a)
}
