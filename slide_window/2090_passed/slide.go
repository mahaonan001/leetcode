package main

import "fmt"

func getAverages(nums []int, k int) []int {
	ans := make([]int, len(nums))
	sum := 0
	i := -k
	for i < len(nums) {
		if i+k < len(nums) {
			sum += nums[i+k]
		}
		if i < k || i > len(nums)-k-1 {
			if i >= 0 {
				ans[i] = -1
			}
			i++
			continue
		}
		ans[i] = sum / (2*k + 1)
		if i >= k {
			sum -= nums[i-k]
		}
		i++
	}
	return ans
}
func main() {
	nums := []int{7, 4, 3, 9, 1, 8, 5, 2, 6}
	k := 3
	fmt.Println(getAverages(nums, k))
}
