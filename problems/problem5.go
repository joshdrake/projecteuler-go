// What is the smallest positive number that is evenly divisible by all of the numbers from 1 to 20?

package main

import "fmt"

func main() {
	i := 0
OUTER:
	for {
		i++
		for j := 1; j <= 20; j++ {
			if i%j != 0 {
				continue OUTER
			}
		}
		break
	}
	fmt.Println(i)
}
