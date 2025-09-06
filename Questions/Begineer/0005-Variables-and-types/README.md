# 0005 - Declare Variables Using `var` and `:=`, and Print Their Types

### Problem Statement

Write a Go program that:

1. Declares variables using both:
   - The `var` keyword
   - The `:=` shorthand
2. Prints the **values** of the variables
3. Prints the **types** of the variables using the `reflect` package

---

### Source Code

```go
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
```

### Explanation

#### Declaring Variables
Go provides two main ways to declare variables:
1. Using var (explicit or implicit typing)
```go
var a int = 10   // Explicit type
var b = 3.14     // Implicit type (Go infers float64)
```


* `var` can be used outside and inside functions.
* You can either specify the type or let Go infer it from the value.

2. Using `:=` (short-hand declaration)
```go
c := true
d := "Gautam Jha"
```

* The := shorthand is only allowed inside functions.
* Go automatically infers the type based on the value.

Prints each variable’s value to the terminal.


### Printing Types
```go
fmt.Println(reflect.TypeOf(a))
```
* The reflect package lets you inspect variable types at runtime.
* Each call returns the type of the given variable.

### Output 
```text
================== Values Of Variables ==================
10
3.14
true
Gautam Jha
================== Values Of Variables ==================

================== Type Of Variables ==================
int
float64
bool
string
================== Type Of Variables ==================
```

