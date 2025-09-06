package main

import "fmt"

func main() {
	name := "Gautam Jha"
	age := 25
	feild := "DevOps"

	fmt.Println(name) // Prints the message and add a newline

	fmt.Printf("My name is %s and my age is %d \n", name, age) // Formatted string, it does not adds a new line so we explicitely have to add it using \n

	FormattedString := fmt.Sprintf("My name is %s, my age is %d and i work in %v feild", name, age, feild) // returns instead of printing to screen

	fmt.Println(FormattedString)
}
