
package main

import (
	"fmt"
	"math"
)

func isPalindrome(s string) bool {
	l := len(s)
	m := int(math.Floor(float64(l / 2)))
	for i := 0; i < m; i++ {
		if s[i] != s[l-1-i] {
			return false
		}
	}

	return true
}

func removeIdx(s string) int {
	l := len(s)
	for i := 0; i < l; i++ {
		j := l - 1 - i
		if s[i] == s[j] {
			continue
		}

		// remove
		if isPalindrome(s[i+1 : j+1]) {
			return i
		} else if isPalindrome(s[i:j]) {
			return j
		}

		return -1
	}

	return -1
}

func printIdx(s string) {
	fmt.Println(removeIdx(s))
}

func main() {
	var t int
	fmt.Scan(&t)

	for i := 0; i < t; i++ {
		var s string
		fmt.Scan(&s)
		printIdx(s)
	}
}