// What is the largest prime factor of the number 600851475143?

package main

import (
	"fmt"
	"../primes"
)

func main() {
	var composite uint64 = 600851475143
	composite_primes := primes.Sieve(composite)
	var factors []uint64
	for _, prime := range composite_primes {
		if composite%prime == 0 {
			factors = append(factors, prime)
		}
	}

	fmt.Println(factors[len(factors)-1])
}
