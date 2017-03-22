package main

import (
	"fmt"
)

func main() {
	//Enter your code here. Read input from STDIN. Print output to STDOUT
	var elementOfArray, temp, result int

	elementOfArray = readNumber()
	for i := 0; i < elementOfArray; i++ {
		temp = readNumber()
		result += temp
	}

	fmt.Println(result)

}

func readNumber() int {
	var n int
	_, err := fmt.Scan(&n)
	if err != nil {
		panic(err)
	}

	return n
}