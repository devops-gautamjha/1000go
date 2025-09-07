## 0014-person-pointer-receiver

### Problem Statement

**Add a method with a pointer receiver to `Person` that mutates it.**

You need to:

* Define a struct `Person` with fields `Name` and `Age`.
* Add methods with **pointer receivers** that modify the struct fields.
* Demonstrate that calling these methods changes the original struct.

---

### Explanation

This program demonstrates how **pointer receivers** in Go methods allow you to **mutate** (change) the original struct instance.

Steps:

1. Define the `Person` struct.
2. Implement methods like `UpdateAge` and `Rename` with pointer receivers `(p *Person)`.
3. These methods update the `Age` and `Name` fields respectively.
4. Create a `Person` instance, print its original state.
5. Call the methods to mutate the struct.
6. Print the updated struct to confirm the changes.

---

### Source Code

```go
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
```

---

### Concept: Pointer Receiver in Methods

* Methods with **pointer receivers** receive the address of the struct (`*Person`).
* This allows the method to **modify the original struct**, not just a copy.
* Pointer receivers are used when:

  * The method needs to modify the receiver.
  * The struct is large and you want to avoid copying.

**Syntax:**

```go
func (p *Person) MethodName() {
    // p points to the original struct, so changes persist
}
```

---

### Output

```
Before update: {Alice 30}
After update: {Alice Smith 35}
```

This output confirms that the methods successfully mutated the original `Person` instance.

