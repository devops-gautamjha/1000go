## 0015-person-stringer-interface

### Problem Statement

**Implement and use an interface (e.g., `Stringer`) for `Person`.**

You need to:

* Define a struct `Person` with fields `Name` and `Age`.
* Implement the `String()` method to satisfy the `fmt.Stringer` interface.
* Print the `Person` struct and observe the custom string output.

---

### Explanation

This program shows how to implement the **`fmt.Stringer` interface** in Go, which requires defining a `String()` method for a type.

Steps:

1. Define the `Person` struct.
2. Implement the `String()` method with a **value receiver** that returns a formatted string describing the `Person`.
3. When printing the `Person` instance with `fmt.Println`, Go automatically calls the `String()` method.
4. The output reflects your custom string representation instead of the default struct format.

---

### Source Code

```go
package main

import (
	"fmt"
)

type Person struct {
	Name string
	Age  int
}

// Implement Stringer interface
func (p Person) String() string {
	return fmt.Sprintf("Person(Name: %s, Age: %d)", p.Name, p.Age)
}

func main() {
	p := Person{Name: "Alice", Age: 30}
	fmt.Println(p) // Automatically calls p.String()
}
```

---

### Concept: `fmt.Stringer` Interface

* The `Stringer` interface in Go is defined as:

  ```go
  type Stringer interface {
      String() string
  }
  ```
* Types that implement this interface can customize how they are printed.
* The `fmt` package’s print functions automatically call the `String()` method if implemented.
* This lets you control how your struct appears in output, logs, or debugging.

---

### Output

```
Person(Name: Alice, Age: 30)
```

This confirms that Go called the custom `String()` method instead of the default struct print.

