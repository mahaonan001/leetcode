package main

import "fmt"

func maxArea(height []int) int {
	result := 0
	l := len(height) / 2
	r := l + 1
	left, right := l-1, r+1
	for left >= 0 && right <= len(height)-1 {
		if height[left] >= height[l] {
			l = left
		}
		if height[right] >= height[r] {
			r = right
		}
		left--
		right++
		result = (r - l) * min(height[l], height[r])
	}
	left, right = l-1, r+1
	for left >= 0 && right <= len(height)-1 {
		tmp := (right - left) * min(height[left], height[right])
		if tmp > result {
			result = tmp
		}
		left--
		right++
	}
	return result
}

// (r-l)*min(r,l)
func main() {
	height := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	fmt.Println(maxArea(height))
}
