// What is the 10,001st prime number?

package main

import (
	"fmt"
	"../primes"
)

func main() {
	var i, j uint64
	i = 1
	for j = 3; i < 10001; j += 2 {
		if primes.IsPrime(j) {
			i++
		}
	}

	fmt.Println(j - 2)
}
