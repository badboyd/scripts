package main

import (
	"fmt"
)

func triColor(arrs []int, mid int) {
	i := 0
	j := 0
	n := len(arrs) - 1
	for j <= n {
		fmt.Println(i, j, n)
		fmt.Println(arrs)
		switch {
		case arrs[j] < mid:
			arrs[i], arrs[j] = arrs[j], arrs[i]
			i++
			j++
		case arrs[j] > mid:
			arrs[j], arrs[n] = arrs[n], arrs[j]
			n--
		default:
			j++
		}
	}
}

func main() {
	a := []int{2, 9, 7, 3, 4, 5, 6, 8, 1, 3, 5, 3}
	triColor(a, 3)
	fmt.Println(a)
}
