package main

import (
	"fmt"
)

func maxSatisfied(customers []int, grumpy []int, minutes int) int {
	cusNum := [2]int{0, 0}
	maxAng := 0

	for k, v := range customers {
		cusNum[grumpy[k]] += v
		if k < minutes-1 {
			continue
		}
		maxAng = max(maxAng, cusNum[1])
		if grumpy[k-minutes+1] == 1 {
			cusNum[1] -= customers[k-minutes+1]
		}
	}
	return cusNum[0] + maxAng
}

func main() {
	fmt.Println(maxSatisfied([]int{1, 0, 1, 2, 1, 1, 7, 5}, []int{0, 1, 0, 1, 0, 1, 0, 1}, 3) == 16)
	fmt.Println(maxSatisfied([]int{10, 1, 7}, []int{0, 0, 0}, 2) == 18)
}
