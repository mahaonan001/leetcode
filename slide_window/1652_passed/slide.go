package main

import (
	"fmt"
)

func decrypt(code []int, k int) []int {
	l := len(code)
	ans := make([]int, l)
	r := k + 1
	if k < 0 {
		r = l
		k = -k
	}
	s := 0
	for _, v := range code[r-k : r] {
		s += v
	}
	for i := range ans {
		ans[i] = s
		s += code[r%l] - code[(r-k)%l]
		r++
	}
	return ans
}

func main() {
	code, k := []int{2, 4, 9, 3}, -2
	fmt.Println(decrypt(code, k))
}
