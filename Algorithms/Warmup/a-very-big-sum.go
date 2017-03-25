package main

import (
	"fmt"
	"math"
)

func main() {
	// Number of size array
	var n int
	var num int
	var sum int

	n = readNumber()

	for i:=1; i <= n; i++ {
		num = readNumber()
		if 0 <= num && num <= int(math.Pow(10,10))  {
			sum += num
		}
	}

	fmt.Println(sum)

}

func readNumber() int {
	var n int
	_, err := fmt.Scan(&n)
	if err != nil {
		panic(err)
	}

	return n
}