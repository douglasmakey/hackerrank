package main

import (
	"fmt"
	"io"
)

func main()  {
	var numbers []int
	var result []int

	for {
		var n int
		_, err := fmt.Scan(&n)
		if err != nil {
			if err == io.EOF {
				break
			} else {
				panic(err)
			}
		}
		numbers = append(numbers, n)
	}

	for _, value := range numbers {
		var sumArray int
		for _, j := range numbers {
			sumArray = sumArray + j
		}

		result = append(result, sumArray - value)
	}

	var smallest, biggest int
	for _, v := range result {
		if v > biggest {
			biggest = v
		}
		if v < biggest {
			smallest = v
		}
	}

	fmt.Println(smallest, biggest)

}