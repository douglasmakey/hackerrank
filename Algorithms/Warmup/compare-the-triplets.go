package main

import (
	"fmt"
)

func main() {

	var quantityRatiing int = 3

	// Init points
	var pointAlice int = 0
	var pointBob int = 0


	// Slice rating
	rating_alice := make([]int, quantityRatiing)
	rating_bob := make([]int, quantityRatiing)

	ReadNumber(rating_alice, 0, quantityRatiing)
	ReadNumber(rating_bob, 0, quantityRatiing)

	pointAlice, pointBob = Compare(rating_alice, rating_bob)

	fmt.Println(pointAlice, pointBob)

}

func ReadNumber(ratings []int, i, n int) {
	if n == 0 {
		return
	}
	if m, err := Scan(&ratings[i]); m != 1 {
		panic(err)
	}
	ReadNumber(ratings, i+1, n-1)
}

func Scan(a *int) (int, error) {
	return fmt.Scan(a)
}

func Compare(rating1, rating2 []int)(int, int) {
	var a int = 0
	var b int = 0
	for key, valueA := range rating1 {
		switch  {
		case valueA > rating2[key]:
			a += 1
		case valueA < rating2[key]:
			b += 1
		}
	}
	return a, b
}
