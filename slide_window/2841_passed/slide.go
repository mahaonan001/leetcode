package main

import "fmt"

func maxSum(nums []int, m int, k int) int64 {
	ans := 0
	res := 0
	sameCount := 0
	aprCount := make(map[int]int)
	for i, v := range nums {
		res += v
		aprCount[v]++
		ct := aprCount[v]
		if ct > 1 {
			sameCount++
		}
		if i < k-1 {
			continue
		}
		if sameCount <= k-m {
			ans = max(ans, res)
		}
		res -= nums[i-k+1]
		if aprCount[nums[i-k+1]] > 1 {
			sameCount--
		}
		aprCount[nums[i-k+1]] -= 1
	}
	return int64(ans)
}
func main() {
	fmt.Println(maxSum([]int{2, 6, 7, 3, 1, 7}, 3, 4))
	fmt.Println(maxSum([]int{5, 9, 9, 2, 4, 5, 4}, 1, 3))
	fmt.Println(maxSum([]int{1, 2, 1, 2, 1, 2, 1}, 3, 3))
	fmt.Println(maxSum([]int{1}, 1, 1))
}
