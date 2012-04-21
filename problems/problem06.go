// Find the difference between the sum of the squares of the first one hundred natural numbers and the square of the sum.

package main

import "fmt"

func main() {
	a, b := 0, 0
	for i := 0; i <= 100; i++ {
		a += i * i
		b += i
	}
	fmt.Println((b * b) - a)
}
