package main

import (
	"fmt"
	"math"
)

func main()  {
	var n int
	a := [][]int{}

	n = readNumber()

	for i:=1; i <= n; i++ {
		var a_t []int
		for j:=1; j <= n; j++ {
			n_t := readNumber()
			a_t = append(a_t, n_t)
		}
		a = append(a, a_t)
	}

	var first int
	var second int

	for idx, row := range  a {
		first += row[idx]
		second += row[n-idx-1]
	}

	var result int

	result = first - second
	fmt.Println(math.Abs(float64(result)))
}


func readNumber() int {
	var n int
	_, err := fmt.Scan(&n)
	if err != nil {
		panic(err)
	}

	return n
}
