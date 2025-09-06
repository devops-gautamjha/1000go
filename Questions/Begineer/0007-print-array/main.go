package main

import "fmt"

func main() {
	fmt.Println("Printing array of 5 Integer")

	var a [5]int
	a[0] = 1
	a[1] = 2
	a[2] = 3
	a[3] = 4
	a[4] = 5

	b := [5]int{1, 2, 3, 4, 5}

	var c = [5]int{}

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

}
