package main

import "fmt"

func main() {
	age := map[string]int{
		"Gautam": 24,
		"Alice":  30,
		"Bob":    25,
	}

	for key, value := range age {
		fmt.Printf("The age of %s is %d\n", key, value)
	}

	// check age of person which doesn't exist in map
	fmt.Println(age["Tom"]) // Will print zero value (in this case zero value of int is 0)

	// Comma Ok idiom

	value, ok := age["Tom"]
	if ok {
		fmt.Println("Age of Tom is: ", value)
	} else {
		fmt.Println("Tom not found")
	}
}
