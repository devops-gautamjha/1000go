# 0002 - Command Line Arguments

###  Problem Statement

Write a Go program that accepts and prints **command-line arguments**.

**Example:**

```bash
go run main.go devops engineer
```
**Output:**
```text
[devops engineer]
```

### Source code 
```go
package main

import (
	"fmt"
	"os"
)

func main() {
	argsWithoutProg := os.Args[1:]

	fmt.Println(argsWithoutProg)
}
```

### Explanation
`import "os"`
- We import the `os` package, which provides access to OS-level arguments, environment variables, and more.
- In this program, we are using it to read the command-line arguments.

`os.Args`
- `os.Args` is a slice of string ([] string).
- It contains all arguments passed to the program, inclusing
    * The progrm name itself (`os.Args[0]`)
    * Any other arguments (`os.Args[1]`, `os.Args[2]`,...)
- Example:
```go
os.Args => ["main", "devops", "engineer"]
```

`argsWithoutProg := os.Args[1:]`
- This line creates a new slice containing only the arguments after the program name.
- `os.Args[1:]` means:
    * "Give me everything from index 1 to the end"


`fmt.Println(argsWithoutProg)`
- This prints the slice of arguments.
- In Go, printing a slice like this will show it with square brackets:
```text
[devops engineer]
```
