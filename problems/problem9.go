// There exists exactly one Pythagorean triplet for which a + b + c = 1000.
// Find the product abc.

package main

import "fmt"

func main() {
	for a := 1; a <= 997; a++ {
		for b := 2; b <= 998; b++ {
			for c := 3; c <= 999; c++ {
				if a + b + c == 1000 {
					if (a*a) + (b*b) == (c*c) {
						fmt.Println(a * b * c)
						return
					}
				}
			}
		}
	}
}