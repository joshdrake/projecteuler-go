// How many routes are there through a 20x20 grid?

package main

import (
	"math/big"
	"fmt"
)

func main() {
	m := int64(20)
	n := int64(20)

	mn_fact := big.NewInt(0)
	m_fact := big.NewInt(0)
	n_fact := big.NewInt(0)
	mn_factprod := big.NewInt(0)

	mn_fact.MulRange(int64(1), (m + n))
	m_fact.MulRange(int64(1), m)
	n_fact.MulRange(int64(1), n)
	mn_factprod.Mul(n_fact, m_fact)

	mn_fact.Div(mn_fact, mn_factprod)

	fmt.Println(mn_fact)
}
