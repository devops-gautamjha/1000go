# 0008 - Slices and Functions

---

###  Problem Statement

Write a Go program to:

- Create slices from arrays
- Append values to slices
- Print slices, and display their `len` and `cap`

---

###  Source Code

```go
package main

import (
	"fmt"
)

func main() {
	// Prime Numbers array
	primeNumber := [6]int{2, 3, 5, 7, 11, 13}

	// Creating slices from the array
	var s []int = primeNumber[1:4] // Elements at index 1, 2, 3 => [3 5 7]
	var s2 []int = primeNumber[:]  // Entire array as slice

	fmt.Println(s)  // [3 5 7]
	fmt.Println(s2) // [2 3 5 7 11 13]

	// Appending to slices
	s = append(s, 17, 19, 23)
	s2 = append(s2, 17, 19, 23)

	fmt.Println(s)  // [3 5 7 17 19 23]
	fmt.Println(s2) // [2 3 5 7 11 13 17 19 23]

	// Length of slices
	fmt.Println("Length of Slice s  :", len(s))
	fmt.Println("Length of Slice s2 :", len(s2))

	// Capacity of slices
	fmt.Println("Capacity of Slice s  :", cap(s))
	fmt.Println("Capacity of Slice s2 :", cap(s2))
}
```

### Explanation
**What is a Slice?**
A slice is a flexible, resizable view into an array.

Unlike arrays, slices don’t own the data — they just refer to it.
```go
primeNumber[1:4]
```
This creates a slice of the array from index 1 up to (but not including) 4 — i.e., `[3 5 7]`.

### Append
Slices can grow using append:
```go
s = append(s, 17, 19, 23)
```
If the slice doesn’t have enough underlying capacity, Go creates a new array in the background and copies the contents over.

### `len()` vs `cap()`
| Function | Meaning                                                   |
| -------- | --------------------------------------------------------- |
| `len()`  | Number of elements currently in the slice                 |
| `cap()`  | Number of elements the slice can hold before reallocating |


### Simple Explanation of cap() in Go
**What is `cap()`?**
`cap()` tells you how much space is available in the underlying array that supports a slice — starting from the slice’s first element.

Think of a slice as a window into an array.

**Example**
```go
arr := [6]int{1, 2, 3, 4, 5, 6}
slice := arr[1:4]
```

- The full array has 6 elements

- `slice` starts at index 1 → the value `2`

- So the slice can see from index 1 to the end of the array

**That means:**

- `len(slice)` = 3 → elements [2, 3, 4]
- `cap(slice)` = 5 → from index 1 to index 5 (total 5 slots)
- `cap(slice)` = (length of underlying array) − (start index of slice)

**When Does cap() Change?**
If you append to a slice and it runs out of space, Go will:

1. Allocate a new array (usually bigger)

2. Copy the old slice values into it

3. Return a new slice pointing to this new array

So now your slice has:

- A new backing array

- A new capacity (larger than before)

### Note:
Capacity growth in Go is not fixed, but generally follows this pattern:

- Double when small

- Slower growth as it gets large

So if your original capacity was 5, Go may grow it to:
- 10, then 12, then 16, etc.

### Output
```text
[3 5 7]
[2 3 5 7 11 13]
[3 5 7 17 19 23]
[2 3 5 7 11 13 17 19 23]
Lenght of Slice s :-  6
Lenght of Slice s2 :-  9
Capacity of slice s :-  12
Capacity of slice s2 :-  12
```

### Summary 
| Concept             | Meaning                                  |
| ------------------- | ---------------------------------------- |
| `slice := arr[1:4]` | Slice from index 1 to 3                  |
| `append(slice, x)`  | Adds value, may create new backing array |
| `len(slice)`        | Number of elements in slice              |
| `cap(slice)`        | Space available in backing array         |
| `cap()` > `len()`   | Normal; Go reserves space for growth     |


