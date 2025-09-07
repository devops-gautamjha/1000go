package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func (p Person) Describe() {
	fmt.Printf("Name: %s, Age: %d\n", p.Name, p.Age)
}

func main() {
	person1 := Person{Name: "Gautam Jha", Age: 24}
	person1.Describe()
}
