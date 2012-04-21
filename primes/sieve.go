package primes

import "math"

func IsPrime(n uint64) bool {
	var i uint64 = 0

	if n < 2 {
		return false
	} else if n == 2 {
		return true
	}

	sqrt_n := uint64((math.Sqrt(float64(n))))
	for i = 3; i <= sqrt_n; i += 2 {
		if n%i == 0 {
			return false
		}
	}

	return true
}

func Sieve(n uint64) []uint64 {
	sqrt_n := uint64((math.Sqrt(float64(n))))
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
