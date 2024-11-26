package main

import "fmt"

func maxFreq(s string, maxLetters int, minSize int, maxSize int) int {
	ans := 0
	leterTimes := make(map[byte]int)
	resultMap := make(map[string]int)
	tmp := ""
	for size := minSize; size <= maxSize; size++ {
		for k := range s {
			tmp += string(s[k])
			leterTimes[s[k]]++
			if k < size-1 {
				continue
			}
			if len(leterTimes) > maxLetters {
				leterTimes[s[k-size+1]]--
				if leterTimes[s[k-size+1]] == 0 {
					delete(leterTimes, s[k-size+1])
				}
				tmp = tmp[1:]
				continue
			}
			resultMap[tmp]++
			ans = max(ans, resultMap[tmp])
			leterTimes[s[k-size+1]]--
			if leterTimes[s[k-size+1]] == 0 {
				delete(leterTimes, s[k-size+1])
			}
			tmp = tmp[1:]
		}
	}
	return ans
}
func main() {
	fmt.Println(maxFreq("cacbbaab", 2, 3, 6))
}
