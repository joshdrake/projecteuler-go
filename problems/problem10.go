// Find the sum of all the primes below two million.

package main

import (
	"fmt"
	"../primes"
)

func main() {
	var sum uint64 = 0
	values := primes.Sieve(2e6 * 2e6)
	for _, value := range values {
		sum += value
	}
	fmt.Println(sum)
}
