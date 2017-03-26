package main

import (
	"fmt"
	"strconv"
)

func main()  {
	var n int
	var a []int


	n = readNumber()

	for i:=1; i <= n; i++ {
		n_t := readNumber()
		a = append(a, n_t)
	}

	var positive []int
	var negative []int
	var zero []int

	for _, v := range a {
		if v > 0 {
			positive = append(positive, v)
		} else if v < 0 {
			negative = append(negative, v)
		} else if v == 0 {
			zero = append(zero, v)
		}
	}

	fmt.Println(FloatToString(0.00000005))




}

func readNumber() int {
	var n int
	_, err := fmt.Scanf("%d", &n)
	if err != nil {
		panic(err)
	}

	return n
}

func FloatToString(input_num float64) string {
	// to convert a float number to a string
	return strconv.FormatFloat(input_num, 'f', 6, 64)
}