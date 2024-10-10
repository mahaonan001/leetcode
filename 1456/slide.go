package main

import (
	"fmt"
)

// 我的方法耗时2.9s
//
//	func maxVowels(s string, k int) int {
//		ans := 0
//		is_u := map[byte]bool{
//			'a': true,
//			'e': true,
//			'i': true,
//			'o': true,
//			'u': true,
//		}
//		l := 0
//		for l+k-1 < len(s) {
//			tmp := 0
//			last_p := 0
//			for i := l; i < l+k; i++ {
//				if is_u[s[i]] {
//					if last_p == 0 && i != l {
//						last_p = i
//					}
//					tmp++
//				}
//			}
//			ans = max(tmp, ans)
//			if ans == k {
//				return ans
//			}
//			if last_p != 0 {
//				l = last_p
//			} else {
//				l += k
//			}
//		}
//		r := len(s) - k
//		tmp := 0
//		for i := r; i < len(s); i++ {
//			if is_u[s[i]] {
//				tmp++
//			}
//		}
//		ans = max(tmp, ans)
//		return ans
//	}
//
// 艾大佬 0.4s
func maxVowels(s string, k int) (ans int) {
	vowel := 0
	for i, in := range s {
		// 1. 进入窗口
		if in == 'a' || in == 'e' || in == 'i' || in == 'o' || in == 'u' {
			vowel++
			if vowel == k {
				return
			}
		}
		if i < k-1 { // 窗口大小不足 k
			continue
		}
		// 2. 更新答案
		ans = max(ans, vowel)
		// 3. 离开窗口
		out := s[i-k+1]
		if out == 'a' || out == 'e' || out == 'i' || out == 'o' || out == 'u' {
			vowel--
		}
	}
	return
}

func main() {
	s, k := "weallloveyou", 7
	fmt.Println(maxVowels(s, k))
	if maxVowels(s, k) == 4 {
		fmt.Println("success")
	} else {
		fmt.Println("ans is wrong")
	}
}
