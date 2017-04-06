package main

import (
	"fmt"
	"sort"
	"math"
)

func main()  {
	// Declare vars
	var n int
	var a []int
	var count int
	var max int

	// Get N
	n = readNumber()

		// Constraints N
		if 1 <= n && n <= int(math.Pow(10, 5)) {

		// Get all Height
		for i:=1; i <= n; i++ {
			n_t := readNumber()

			// Constraints Height
			if 1 <= n_t && n_t <= int(math.Pow(10,7)) {
				a = append(a, n_t)
			}
		}
	}

	// Sort array
	sort.Ints(a)

	// Get max
	max = a[len(a)-1]

	for i := 0; i < n; i++ {
		if a[i] == max {
			count++;
		}
	}

	fmt.Println(count)

}


func readNumber() int {
	var n int
	_, err := fmt.Scan(&n)
	if err != nil {
		panic(err)
	}

	return n
}

