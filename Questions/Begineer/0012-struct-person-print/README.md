# 0012-struct-person-print

### Problem Statement

**Define a basic struct `Person` with `Name` and `Age` and print it.**

In this task, you are required to:

* Define a struct type called `Person` with fields `Name` (string) and `Age` (int).
* Create instances of this struct using different methods.
* Print the struct values to the console.

---

### Explanation

This program demonstrates how to:

* Define a basic struct in Go.
* Initialize struct values in two different ways:

  * By assigning fields one by one.
  * By using a struct literal.
* Print struct values using both `fmt.Println` (which prints the struct in `{}` format) and `fmt.Printf` for a more formatted output.

---

### Source Code

```go
package main

import "fmt"

func main() {

	type Person struct {
		Name string
		Age  int
	}

	// Initializing struct using field assignment
	var person1 Person
	person1.Name = "Gautam Jha"
	person1.Age = 24

	// Initializing struct using struct literal
	person2 := Person{Name: "Hardik", Age: 25}

	// Printing entire struct
	fmt.Println(person1)

	// Printing struct fields individually with formatting
	fmt.Printf("The name of person 2 is %s and his age is %d\n", person2.Name, person2.Age)
}
```

---

### Concept: Structs in Go

* A **struct** is a composite type in Go that groups together variables (called fields) under one name.
* Structs are used to model real-world entities with multiple attributes.
* Fields can be accessed using the dot `.` notation.

**Example Syntax:**

```go
type Person struct {
	Name string
	Age  int
}
```

Structs can be initialized in two ways:

1. **By assigning fields individually**:

   ```go
   var p Person
   p.Name = "Alice"
   p.Age = 30
   ```
2. **Using a struct literal**:

   ```go
   p := Person{Name: "Bob", Age: 25}
   ```

---

###  Output

```
{Gautam Jha 24}
The name of person 2 is Hardik and his age is 25
```

> Note: When printing a struct using `fmt.Println`, Go displays the values in `{}` format in field order.
