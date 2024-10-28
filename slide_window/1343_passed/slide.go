package main

import "fmt"

func numOfSubarrays(arr []int, k int, threshold int) int {
	ans, sum := 0, 0
	for i := range arr {
		sum += arr[i]
		if i < k-1 {
			continue
		}
		if float64(sum)/float64(k) >= float64(threshold) {
			ans++
		}
		sum -= arr[i-k+1]
	}
	return ans
}

func main() {
	arr := []int{2, 2, 2, 2, 5, 5, 5, 8}
	k, threshold := 3, 4
	fmt.Println(numOfSubarrays(arr, k, threshold))
}
