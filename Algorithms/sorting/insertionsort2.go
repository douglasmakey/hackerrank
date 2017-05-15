package main

import "fmt"

func swap(a []int, indexA, indexB int) {
	var tmp int = a[indexA]
	a[indexA] = a[indexB]
	a[indexB] = tmp
}

func printArr(a []int) {
	for i := 0; i < len(a); i++ {
		fmt.Printf("%v ", a[i])
	}
	fmt.Println()
}

func insertionSort(a []int) {
	for i := 1; i < len(a); i++ {
		for j := i; j > 0 && a[j] < a[j - 1]; j-- {
			swap(a, j, j - 1)
		}
		printArr(a)
	}

}

func main() {
	var a []int
	var t int
	fmt.Scan(&t)

	for i := 0; i < t; i++ {
		var n int
		fmt.Scan(&n)
		a = append(a, n)
	}

	insertionSort(a)

}