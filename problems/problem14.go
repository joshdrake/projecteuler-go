// Which starting number, under one million, produces the longest Collatz chain?

package main

import "fmt"

func main() {
	var max, number, i uint64
	for ; i < 1e6; i++ {
		length := chain_length(i)
		if length > max {
			max = length
			number = i
		}
	}
	fmt.Println(number)
}

func chain_length(n uint64) uint64 {
	length := uint64(0)
	for n > 1 {
		if n%2 == 0 {
			n = n / 2
		} else {
			n = (n * 3) + 1
		}
		length++
	}
	return length
}
