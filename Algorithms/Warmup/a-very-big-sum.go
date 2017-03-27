package main

import (
	"fmt"
	"io"
	"math"
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

		if n >= 1 && n <= int(math.Pow(10,9)) {
			numbers = append(numbers, n)
		}
	}

	for _, value := range numbers {
		var sumArray int
		for _, j := range numbers {
			sumArray = sumArray + j
		}

		result = append(result, sumArray - value)
	}


	smallest, biggest := result[0], result[0]
	for _, v := range result {
		if v > biggest {
			biggest = v
		}
		if v < smallest {
			smallest = v
		}
	}

	fmt.Println(smallest, biggest)

}