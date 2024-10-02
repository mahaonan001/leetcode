package main

import (
	"fmt"
)

const (
	MaxPost    = (1 << 31) - 1
	MiniNeg    = -1 << 31
	OutMiniNeg = -1 << 62
	OutMaxPost = 1<<62 - 1
)

func add2(result *int64, times *int, n byte) {
	*result *= 10
	*result += int64(n - 48)
	*times++
}

func myAtoi(s string) int {
	var result int64 = 0
	times := 0
	valid := false
	neg := false
	before := false
	outTime := 0
	last := true
	for i := 0; i < len(s); i++ {
		if !valid && string(s[i]) == " " {
			continue
		} else if valid && string(s[i]) == " " {
			break
		}
		if valid {
			if string(s[i]) < "0" || string(s[i]) > "9" {
				break
			} else {
				// if string(s[i]) == "0" && i != len(s)-1 {
				// 	continue
				// }
				if string(s[i]) != "0" {
					before = true
				}
				if before {
					add2(&result, &times, s[i])
					if result < 0 && last {
						outTime++
						last = false
					}
					if result > 0 && !last {
						outTime++
						last = true
					}
				}
				continue
			}
		}
		if !valid && (string(s[i]) == "-" || (string(s[i]) >= "0" && string(s[i]) <= "9")) {

			if string(s[i]) == "-" {
				neg = true
			} else {
				if string(s[i]) != "0" {
					before = true
				}
				add2(&result, &times, s[i])
			}
			valid = true
			continue
		} else if string(s[i]) == "+" {
			valid = true
			continue
		} else {
			break
		}

	}
	if neg {
		result = -result
	}
	fmt.Println(result)
	fmt.Println(outTime)
	fmt.Println(neg)
	if outTime%2 == 1 && result > 0 {
		if !neg {
			return MaxPost
		}
		return MiniNeg
	}
	if outTime%2 == 1 && result < 0 {
		if neg {
			return MiniNeg
		}
		return MaxPost
	}
	if outTime%2 == 0 && outTime > 0 && result > 0 {
		if !neg {
			return MaxPost
		}
		return MiniNeg
	}
	if outTime%2 == 0 && outTime > 0 && result < 0 {
		if neg {
			return MiniNeg
		}
		return MaxPost
	}
	if result < MiniNeg {
		if (result >> 1) <= OutMiniNeg {
			return MaxPost
		}
		return MiniNeg
	} else if result > MaxPost {
		if (result >> 1) >= OutMaxPost {
			return MiniNeg
		}
		return MaxPost
	}
	return int(result)
}

func main() {
	s := "18446744073709551617"
	fmt.Println(myAtoi(s))
}
