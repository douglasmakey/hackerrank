package main

import (
	"fmt"
	"sort"
)

func main()  {

	// Number of hurdles
	var n int

	// Maximum height he can jump
	var k int

	n = readNumber()
	k = readNumber()

	var height []int

	for i:=1; i <= n; i++ {
		height = append(height, readNumber())
	}

	// Sort height
	sort.Ints(height)

	max_height := height[len(height)-1]

	max_jump := k

	if max_jump >= max_height {
		fmt.Println(0)
	} else {
		fmt.Println(max_height - max_jump)
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
