package main

import (
	"fmt"
)

func Judge(s, p byte) bool {
	return s == p || p == byte('.')
}

func MeetPoint(p []byte, node byte, index int) {
	p[index] = node
}

func DelStar(p []byte, index int) {
	left := p[:index]
	right := p[index+1:]
	p = append(left, right...)
}

func isMatch(s string, p string) bool {
	byteS := []byte(s)
	byteP := []byte(p)
	index := 0
	for k := range byteS {
		fmt.Println(string(byteS[k]), string(byteP[index]))
		if byteS[k] == byteP[index] {
			index++
		} else if byteP[index] == byte('.') {
			index++
		}
		if byteP[index] == byte('*') {
			if byteP[index-1] != byteS[k] && byteP[index-1] != byte('.') {
				index++
			}
			if byteS[k] == byteP[index] {
				index++
			} else if byteP[index] == byte('.') {
				index++
			}
		} else if byteP[index] != byteS[k] {
			break
		}

	}
	fmt.Println(string(byteP))
	return true
}

func main() {
	x := "mississippi"
	fmt.Println(isMatch(x, "mis*is*ip*."))
}
