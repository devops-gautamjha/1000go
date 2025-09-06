package main

import "fmt"

func main() {
	primeNumber := [6]int{2, 3, 5, 7, 11, 13}

	// Creating a slice
	var s []int = primeNumber[:]

	s = append(s, 17, 19, 23)

	for i, v := range s {
		fmt.Printf("Value at index %d is %d \n", i, v)
	}

}
