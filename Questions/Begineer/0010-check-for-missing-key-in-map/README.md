# 0010 - Create and Use a Map; Check for Missing Keys

---

### Problem Statement

Write a Go program to:

- Create a `map[string]int`
- Access values by keys
- Handle missing keys using the **comma-ok idiom**

---

### Source Code

```go
package main

import "fmt"

func main() {
	age := map[string]int{
		"Gautam": 24,
		"Alice":  30,
		"Bob":    25,
	}

	for key, value := range age {
		fmt.Printf("The age of %s is %d\n", key, value)
	}

	// Access a missing key
	fmt.Println(age["Tom"]) // Prints 0 (zero value of int)

	// Use comma-ok idiom to check if key exists
	value, ok := age["Tom"]
	if ok {
		fmt.Println("Age of Tom is:", value)
	} else {
		fmt.Println("Tom not found")
	}
}
```

### What is a Map in Go?

A map in Go is a built-in data structure that stores key-value pairs.
It's similar to:
- `dict` in Python
- `object` in JavaScript (used like a dictionary)
- `HashMap` in Java

### Map Syntax
```go 
map[KeyType]ValueType
```

### Example 
```go
age := map[string]int{
    "Alice": 30,
    "Bob": 25,
}
```
- Keys are of type `string`
- Values are of type `int`
- You access values using the key: `age["Alice"]` → `30`

### Missing Keys & Zero Values
When you try to access a key that does not exist, Go does not throw an error.
Instead, it returns the zero value of the value type.
```go
fmt.Println(age["Tom"]) // Prints 0
```

### Comma-Ok Idiom

To check if a key exists, use this Go pattern:
```go
value, ok := map[key]
```
- `value` is the result

- `ok` is a boolean:

- `true` → key was found

- `false` → key not found

```go
value, ok := age["Tom"]
if ok {
    fmt.Println("Tom's age is:", value)
} else {
    fmt.Println("Tom not found")
}
```
This is the safe and recommended way to check for missing keys in Go.


### Output
```text
The age of Gautam is 24
The age of Alice is 30
The age of Bob is 25
0
Tom not found
```

### Summary

| Concept      | Example                              |
| ------------ | ------------------------------------ |
| Create a map | `age := map[string]int{ "Bob": 25 }` |
| Access value | `age["Bob"]`                         |
| Missing key  | Returns zero value (e.g., `0`)       |
| Safe check   | `value, ok := age["Tom"]`            |
