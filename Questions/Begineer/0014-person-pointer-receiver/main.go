package main

import (
	"fmt"
)

// Step 1: Define the Person type
type Person struct {
	Name string
	Age  int
}

// Step 2: Method with a pointer receiver that mutates the struct
func (p *Person) UpdateAge(newAge int) {
	p.Age = newAge
}

// Another example: change name
func (p *Person) Rename(newName string) {
	p.Name = newName
}

func main() {
	// Step 3: Create a Person instance
	p := Person{Name: "Alice", Age: 30}

	// Show original
	fmt.Println("Before update:", p)

	// Step 4: Call pointer receiver methods to mutate
	p.UpdateAge(35)
	p.Rename("Alice Smith")

	// Show mutated
	fmt.Println("After update:", p)
}
