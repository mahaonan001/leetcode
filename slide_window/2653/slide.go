package main

import (
	"fmt"
)

func getSubarrayBeauty(nums []int, k int, x int) []int {
	ans := make([]int, 0)
	tmp := make([]int, 0)
	for i, v := range nums {
		tmp = append(tmp, v)
		if i < k-1 {
			continue
		}
		after := make([]int, 50)

		for _, v := range nums {
			if v < 0 {
				after[50+v]++
			}
		}
		
		if after[x-1] < 0 {
			ans = append(ans, after[x-1])
		} else {
			ans = append(ans, 0)
		}
		tmp = tmp[1:]
	}
	return ans
}
func main() {
	fmt.Println(getSubarrayBeauty([]int{-3, 1, 2, -3, 0, -3}, 2, 1))
}
