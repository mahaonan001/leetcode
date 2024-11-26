package main

import "fmt"

func resInit(size int, nums []int) (res int) {
	res = 0
	for i := len(nums) - size + 1; i < len(nums); i++ {
		res += nums[i]
	}
	return res
}

func getNumber(index int, nums []int) int {
	if index < 0 {
		index = len(nums) - index
	}
	if index >= len(nums) {
		index = index - len(nums)
	}
	return nums[index]
}

func minSwaps(nums []int) int {
	ans := 0
	size := 0
	for k := range nums {
		size += nums[k]
	}
	res := resInit(size, nums)
	fmt.Println(res)
	for k := range nums {
		res += nums[k]
		if k < size {
			continue
		}
		res -= getNumber(k-size+1, nums)
		ans = max(ans, res)
	}
	return size - ans
}

// 0 1 2 3 4 5 6
// 1 5 7 5 4 3 1

func main() {
	fmt.Println(minSwaps([]int{0, 1, 1, 1, 0, 0, 1, 1, 0}))
}
