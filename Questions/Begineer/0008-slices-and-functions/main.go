package main

import (
	"fmt"
)

func main() {

	// Prime Numbers array
	primeNumber := [6]int{2, 3, 5, 7, 11, 13}

	// Creating a slice
	var s []int = primeNumber[1:4] // This selects a half-open range of array which includes the first element, but excludes the last one.

	var s2 []int = primeNumber[:] // This includes all elements of array
	fmt.Println(s)                // Printing the slice
	fmt.Println(s2)
	// append to slice

	s = append(s, 17, 19, 23)
	s2 = append(s2, 17, 19, 23)

	fmt.Println(s)
	fmt.Println(s2)

	// Length of slice

	fmt.Println("Lenght of Slice s :- ", len(s))
	fmt.Println("Lenght of Slice s2 :- ", len(s2))

	// Capacity of slice

	fmt.Println("Capacity of slice s :- ", cap(s))
	fmt.Println("Capacity of slice s2 :- ", cap(s2))

}
