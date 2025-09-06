# 0004 - Demonstrate `fmt.Println`, `fmt.Printf`, and `fmt.Sprintf`

### Problem Statement

Write a Go program to demonstrate the **differences between**:

- `fmt.Println()`
- `fmt.Printf()`
- `fmt.Sprintf()`

Use example variables to show how each function works and how the output differs.

---

### Source Code

```go
package main

import "fmt"

func main() {
	name := "Gautam Jha"
	age := 25
	feild := "DevOps"

	fmt.Println(name) // Prints the message and adds a newline

	fmt.Printf("My name is %s and my age is %d \n", name, age) // Formatted string, doesn't add newline unless you specify \n

	FormattedString := fmt.Sprintf("My name is %s, my age is %d and i work in %v feild", name, age, feild) // Returns formatted string

	fmt.Println(FormattedString)
}
```

### Explanation — Line by Line
`name := "Gautam Jha"`
Defines a string variable with your name.

`age := 25`
Defines an integer variable representing your age.

`feild := "DevOps"`
(You likely meant field but this still works.) Defines your work area.

### Print Statements Explained
`fmt.Println(name)`
- Prints the value of name with a newline at the end.
- Output:
```text
Gautam Jha
```

`fmt.Printf(...)`
```go
fmt.Printf("My name is %s and my age is %d \n", name, age)
```

- Used for formatted printing.

- `%s` is replaced with the string value of `name`
- `%d` is replaced with the integer value of `age`
- You must manually add a newline (`\n`) if you want the output to break.
- Output:
```text
My name is Gautam Jha and my age is 25 
```

`fmt.Sprintf(...)`
```go
FormattedString := fmt.Sprintf("My name is %s, my age is %d and i work in %v feild", name, age, feild)
```
- Works just like Printf, but instead of printing, it returns the formatted string.
- You can store this string in a variable for later use.
- `%v` is a generic formatter (can handle many types, like `string`, `int`, `etc`.)

Then:
```go
fmt.Println(FormattedString)
```

### Summary of Differences
Prints the returned string.
- Output:
```text
My name is Gautam Jha, my age is 25 and i work in DevOps feild
```

| Function        | What it does                          | Adds newline?      | Returns string? | Use case                                   |
| --------------- | ------------------------------------- | ------------------ | --------------- | ------------------------------------------ |
| `fmt.Println()` | Prints values with spaces             | ✅ Yes              | ❌ No            | Simple printing                            |
| `fmt.Printf()`  | Prints formatted string to console    | ❌ No (unless `\n`) | ❌ No            | Precise, styled output                     |
| `fmt.Sprintf()` | Returns a formatted string (no print) | ❌ No               | ✅ Yes           | Store and use string later (e.g., logging) |


### How to Run
```bash
go run main.go
```

### Output
```text
Gautam Jha
My name is Gautam Jha and my age is 25 
My name is Gautam Jha, my age is 25 and i work in DevOps feild
```

### Real-World DevOps Use Cases
- `Printf`: Useful for progress logs or status messages in CLI tools

- `Sprintf`: When preparing structured messages for logs, alerts, or file writing

- `Println`: Quick, simple console output (debugging, test prints)

