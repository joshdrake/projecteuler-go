// By considering the terms in the Fibonacci sequence whose values
// do not exceed four million, find the sum of the even-valued terms.

package main

import "fmt"

func main() {
	sum := 0
	for a, b := 1, 2; b < 4e6; {
		if b%2 == 0 {
			sum += b
		}
		a, b = b, b+a
	}
	fmt.Println(sum)
}
