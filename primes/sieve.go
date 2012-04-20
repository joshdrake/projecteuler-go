package primes

import (
	"math"
)

func Sieve(n uint64) []uint64 {
	sqrt_n := uint64(math.Floor(math.Sqrt(float64(n))))
	var sieve = make([]bool, sqrt_n+1)
	var i, j uint64
	for i = 0; i <= sqrt_n; i++ {
		sieve[i] = true
	}
	for i = 2; i <= sqrt_n; i++ {
		if sieve[i] {
			for j = 2 * i; j <= sqrt_n; j += i {
				sieve[j] = false
			}
		}
	}

	var primes []uint64
	for i = 2; i <= sqrt_n; i++ {
		if sieve[i] {
			primes = append(primes, i)
		}
	}
	return primes
}

