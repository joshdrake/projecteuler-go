package numbers

import "../primes"

func Factors(n uint64) []uint64 {
	if n == 0 || n == 1 {
		factors := []uint64{n}
		return factors
	}

	prime_factors := primes.Sieve(n)
	trial := make(map[uint64]bool)
	trial[1] = true
	trial[n] = true
	for _, factor := range prime_factors {
		if n%factor == 0 {
			trial[factor] = true
			for _, f := range Factors(n / factor) {
				trial[f] = true
			}
		}
	}

	var factors []uint64

	for factor, trial := range trial {
		if trial {
			factors = append(factors, factor)
		}
	}

	return factors
}
