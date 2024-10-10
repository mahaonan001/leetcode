package main

import "fmt"

func maxArea(height []int) int {
	l := (len(height) - 1) / 2
	r := l + 1
	result := (r - l) * min(height[l], height[r])
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
	left, right = l, r+1
	for left >= 0 && right < len(height) {
		tmp := (right - left) * min(height[left], height[right])
		if tmp > result {
			result = tmp
		}
		right++
	}
	left, right = l-1, r
	for left >= 0 && right <= len(height)-1 {
		tmp := (right - left) * min(height[left], height[right])
		if tmp > result {
			result = tmp
		}
		left--
	}
	return result
}

// (r-l)*min(r,l)
func main() {
	height := []int{8, 7, 2, 1}
	fmt.Println(maxArea(height))
}
