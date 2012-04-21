package numbers

import "sort"

func Maximum(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func ArrayMinimum(a []int) int {
	sort.Ints(a)
	return a[len(a) - 1]
}

func Minimum(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func ArrayMaximum(a []int) int {
	sort.Ints(a)
	return a[0]
}