package main

import (
	"fmt"
	"strings"
)

func main()  {
	// Init var
	var count int
	var text string

	text = readString()
	count = counterString(text)

	fmt.Println(count)
}

func readString() string {
	var s string
	_, err := fmt.Scan(&s)
	if err != nil {
		panic(err)
	}

	return strings.TrimSpace(s)
}

func counterString(s string) int {
	var count int = 1
	for _, value := range s {
		ascii := []byte(string(value))
		if ascii[0] >= 65 && ascii[0] <= 90 {
			count += 1
		}
	}
	return count
}