package main

import "fmt"

func minimumRecolors(blocks string, k int) int {
	ans := 0
	num := 0
	for i := range blocks {
		if blocks[i] == 'B' {
			num++
		}
		if i < k-1 {
			continue
		}
		if num == k {
			return 0
		}
		ans = max(ans, num)
		if blocks[i-k+1] == 'B' {
			num--
		}
	}
	return k - ans
}
func main() {
	blocks := "WBWBBBW"
	k := 2
	fmt.Println(minimumRecolors(blocks, k))
}
