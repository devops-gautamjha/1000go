# 0009 - Iterate Over a Slice with `for` and `range`

---

### 🧩 Problem Statement

Write a Go program to:

- Create a slice of integers
- Use a `for` loop with `range` to iterate over the slice
- Print both the index and value at each position

---

### 📄 Source Code

```go
package main

import "fmt"

func main() {
	// Step 1: Define an array
	primeNumber := [6]int{2, 3, 5, 7, 11, 13}

	// Step 2: Create a slice from the array
	var s []int = primeNumber[:]

	// Step 3: Append more values to the slice
	s = append(s, 17, 19, 23)

	// Step 4: Use for + range to iterate over slice
	for i, v := range s {
		fmt.Printf("Value at index %d is %d \n", i, v)
	}
}
```

### Explanation

**What is range?**
In Go, range is used with for loops to iterate over arrays, slices, maps, strings, or channels.

When used with slices or arrays:

```go
for index, value := range slice {
    // your logic
}
```

- index → the current index (0, 1, 2, ...)
- value → the value at that index

### Step-by-Step
1. You start with an array of prime numbers
2. Create a slice of the full array:
```go
s := primeNumber[:]
```
3. Append more primes:
```go
s = append(s, 17, 19, 23)
```
4. Use `for i, v := range s` to:
- Get the index `i`
- Get the value `v`
- Print them both

### Output

```text
Value at index 0 is 2 
Value at index 1 is 3 
Value at index 2 is 5 
Value at index 3 is 7 
Value at index 4 is 11 
Value at index 5 is 13 
Value at index 6 is 17 
Value at index 7 is 19 
Value at index 8 is 23
```

### Common Variations

- You can ignore the index if you only need the value:
```go
for _, v := range s {
    fmt.Println(v)
}
```
- Or ignore the value:
```go
for i := range s {
    fmt.Println(i)
}
```

### Summary
| Concept         | Explanation                                   |
| --------------- | --------------------------------------------- |
| `for ... range` | Loop through each element in a slice or array |
| `i`             | Current index                                 |
| `v`             | Value at that index                           |
| `append()`      | Adds values to a slice                        |


