package main

import (
	"fmt"
	"strings"
)


func main() {
	var n int

	n = readNumber()

	for i:=1; i <= n; i++ {
		hash_count := strings.Repeat("#", i)
		difference := n - i
		space := strings.Repeat(" ", difference)
		output := space + hash_count
		fmt.Println(output)
	}
}



func readNumber() int {
	var n int
	_, err := fmt.Scan(&n)
	if err != nil {
		panic(err)
	}

	return n
}
