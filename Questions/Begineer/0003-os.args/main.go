package main

import (
	"fmt"
	"os"
)

func main() {
	argsWithProg := os.Args

	nameOfProgram := argsWithProg[0]
	numberOfArgs := len(argsWithProg[1:])

	fmt.Println("Name of Program is :- ", nameOfProgram)

	fmt.Println("Number of arguments are :- ", numberOfArgs)

}
