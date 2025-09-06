# 0007 - Print Array of 5 Integers

---

### Problem Statement

Write a Go program to:

- Declare an array of 5 integers
- Initialize it in different ways
- Print the contents of each array

---

### Source Code

```go
package main

import "fmt"

func main() {
	fmt.Println("Printing array of 5 Integer")

	// Method 1: Declare and assign one by one
	var a [5]int
	a[0] = 1
	a[1] = 2
	a[2] = 3
	a[3] = 4
	a[4] = 5

	// Method 2: Declare and initialize in one line
	b := [5]int{1, 2, 3, 4, 5}

	// Method 3: Declare empty array (defaults to all 0)
	var c = [5]int{}

	// Print all arrays
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}
```
### Explanation
Go arrays are fixed-size, indexed collections of elements with the same type.

```go
var a [5]int
```

This creates an array of 5 integers. All values are initialized to `0` by default.

### Different Ways to Work With Arrays
| Method                  | Code                         | Explanation                        |
| ----------------------- | ---------------------------- | ---------------------------------- |
| Manual assignment       | `a[0] = 1`                   | You assign each element one-by-one |
| One-line initialization | `b := [5]int{1, 2, 3, 4, 5}` | Set values all at once             |
| Empty array             | `var c [5]int`               | Defaults to `[0 0 0 0 0]`          |


### Output
```text
Printing array of 5 Integer
[1 2 3 4 5]
[1 2 3 4 5]
[0 0 0 0 0]
```

### Quick Notes
- Array indices start at 0
- Array length is fixed — you can't grow or shrink it
- If you need resizable arrays, use slices

### Summary
| Concept                | Example                      |
| ---------------------- | ---------------------------- |
| Declare empty array    | `var a [5]int`               |
| Assign values by index | `a[0] = 10`                  |
| Initialize with values | `b := [5]int{1, 2, 3, 4, 5}` |
| Print entire array     | `fmt.Println(a)`             |

