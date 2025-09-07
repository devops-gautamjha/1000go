# 0013-person-value-receiver

### Problem Statement

**Add a method with a value receiver to `Person`.**

You are required to:

* Define a struct type `Person` with fields like `Name` and `Age`.
* Implement a method (e.g., `Describe`) **using a value receiver**.
* Call the method from a `Person` instance and print its output.

---

### Explanation

This program introduces the concept of **methods with value receivers** in Go.

Steps:

1. A `Person` struct is defined with `Name` and `Age` fields.
2. A method `Describe()` is added to the `Person` struct using a **value receiver** `(p Person)`.
3. This method prints out the details of the `Person`.
4. The method is then called on an instance of the struct.

---

### Source Code

```go
package main

import "fmt"

// Define the struct
type Person struct {
	Name string
	Age  int
}

// Method with a value receiver
func (p Person) Describe() {
	fmt.Printf("Name: %s, Age: %d\n", p.Name, p.Age)
}

func main() {
	person1 := Person{Name: "Gautam Jha", Age: 24}
	person1.Describe()
}
```

---

### Concept: Value Receiver in Methods

* In Go, a **method** is a function with a **receiver**.
* A **value receiver** means the method receives a **copy** of the struct.
* Any changes made inside the method **do not affect** the original object.

**Syntax:**

```go
func (p Person) MethodName() {
    // p is a copy of the original struct
}
```

Use value receivers when:

* The method doesn't need to modify the receiver.
* The receiver is small and copying is cheap.

---

### Output

```
Name: Gautam Jha, Age: 24
```

This confirms that the `Describe()` method correctly accessed the struct’s fields and printed them.

