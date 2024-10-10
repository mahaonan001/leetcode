package main

import (
	"fmt"
)

func decrypt(code []int, k int) []int {
	l := len(code)
	ans := make([]int, l)
	switch {
	case k == 0:
		return ans
	case k > 0:
		return ans
	case k < 0:
		return ans
	}
	return ans
}

func main() {
	code, k := []int{2, 4, 9, 3}, -2
	fmt.Println(decrypt(code, k))
}
