package main

import "fmt"

func findMaxAverage(nums []int, k int) float64 {
	ans := -999999
	tmp := 0
	for i := range nums {
		tmp += nums[i] //slide  in
		if i-k+1 < 0 {
			continue // judge
		}
		out := i - k + 1
		ans = max(ans, tmp)
		tmp -= nums[out] // come out
	}
	return float64(ans) / float64(k)
}
func main() {
	nums := []int{-1}
	k := 1
	fmt.Println(findMaxAverage(nums, k))
}
