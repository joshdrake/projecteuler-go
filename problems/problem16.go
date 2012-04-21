// What is the sum of the digits of the number 21000?

package main

import (
	"math/big"
	"fmt"
	"strconv"
)

func main() {
	n := big.NewInt(2)
	sum := 0

	n.Exp(n, big.NewInt(1000), nil)

	for _, m := range n.String() {
		m, _ := strconv.Atoi(string(m))
		sum += m
	}

	fmt.Println(sum)
}
