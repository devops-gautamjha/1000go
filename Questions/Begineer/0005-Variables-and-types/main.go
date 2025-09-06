package main

import (
	"fmt"
	"reflect"
)

func main() {
	// Variable declaration using var keyword
	var a int = 10 // Specify type explicitly
	var b = 3.14   // Let Go infer the type (float64)

	// Variable declaration using :=
	c := true
	d := "Gautam Jha"

	// Printing values of variables
	fmt.Println("================== Values Of Variables ==================")
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println("================== Values Of Variables ==================")

	// Blank line for readability
	fmt.Println()

	// Printing types of variables
	fmt.Println("================== Type Of Variables ==================")
	fmt.Println(reflect.TypeOf(a))
	fmt.Println(reflect.TypeOf(b))
	fmt.Println(reflect.TypeOf(c))
	fmt.Println(reflect.TypeOf(d))
	fmt.Println("================== Type Of Variables ==================")
}
