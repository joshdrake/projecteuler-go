// What is the value of the first triangle number to have over five hundred divisors?

package main

import (
	"fmt"
	"../numbers"
)

func main() {
	triangle := uint64(1)
	for i := uint64(2); ; i++ {
		factors := numbers.Factors(triangle)
		if len(factors) > 500 {
			fmt.Println(triangle)
			return
		}
		triangle += i
	}
}
