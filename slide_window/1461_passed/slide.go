package main

import "fmt"

func hasAllCodes(s string, k int) bool {
	byteBook := make(map[string]bool)
	for i := 0; i < len(s)-k+1; i++ {
		byteBook[s[i:k+i]] = true
	}
	return len(byteBook) == 1<<k
}

func main() {
	fmt.Println(hasAllCodes("00110", 2))
}
