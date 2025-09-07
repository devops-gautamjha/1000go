package main

import (
	"fmt"
)

type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("Person(Name: %s, Age: %d)", p.Name, p.Age)
}

func main() {
	p := Person{Name: "Alice", Age: 30}
	fmt.Println(p) // Automatically calls p.String()
}
