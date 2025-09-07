package main

import "fmt"

func main() {

	type Person struct {
		Name string
		Age  int
	}

	var person1 Person
	person1.Name = "Gautam Jha"
	person1.Age = 24

	person2 := Person{Name: "Hardik", Age: 25}

	fmt.Println(person1)

	fmt.Printf("The name of person 2 is %s and his age is %d\n", person2.Name, person2.Age)

}
