# 0003 - Print Program Name and Number of Arguments

### Problem Statement

Write a Go program that:

1. Prints the **name of the program**
2. Prints the **number of arguments passed to it (excluding the program name)**

---

### Source Code

```go
package main

import (
	"fmt"
	"os"
)

func main() {
	argsWithProg := os.Args

	nameOfProgram := argsWithProg[0]
	numberOfArgs := len(argsWithProg[1:])

	fmt.Println("Name of Program is :- ", nameOfProgram)

	fmt.Println("Number of arguments are :- ", numberOfArgs)

}
```

### Explanation 
`os.Args`
- `os.Args` is a slice (`[]string`) that holds:
    * The program name at index 0
    * All command-line arguments at index 1 onward
- `argsWithProg := os.Args`
    * This just stores the full slice (including the program name) into a variable for clarity.


`nameOfProgram := argsWithProg[0]`
- This retrieves the name of the Go program (or script) that is being executed.
- On most systems, it will be something like:
    * `main` (if run with `go run main.go`)
- Or `./myprogram` (if built and executed directly)

`numberOfArgs := len(argsWithProg[1:])`
- `argsWithProg[1:]` gives us all the actual arguments, excluding the program name.
- `len(...)` returns the count of those arguments.


`fmt.Println(...)`
- First prints the program name.
- Then prints how many arguments were passed in total.

### How to run 
```bash
go run main.go devops tools golang
```

### Example Output 
```text
Name of Program is :-  /tmp/go-build1234/b001/exe/main
Number of arguments are :-  3
```

