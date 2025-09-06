# 0001 - Hello, World!

###  Problem Statement

Write a simple Go program that prints **"Hello, World!"** to the console.

This is the traditional first program in any programming language and is used to demonstrate the basic structure of a program.

---

###  Source Code

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, World!")
}
```

### Explanation 

`packge main` 
* Every Go program starts with a package declaration.
* `main` is a special package name.
    - It is used to tell go compiler:
        > "This is the entrypoint of a standalone executable program."
* Without `package main`, your code cannot run as a program.

`import "fmt"`
* This line tells Go to import the `fmt` package, which stands for formatting.
* `fmt` is part of Go's standard library and provides functions for:
    - Printing to the console (`fmt.Println`, `fmt.Print`)
    - Scanning input from the users
* We're using it here to print text to the screen.

`func main() { ... }`
- This is the main function - the starting point of the execution in Go.
- Go always starts running the code from `func main()` when the program is launched.
- Think of it as the "entry door" to your application.

`fmt.Println("Hello, World")`
- This line calls the `Println` function from the `fmt` package.
- `Println` stands for Print Line, which:
    * Prints the given string to the console
    * Automatically adds a newline at the end
- So it prints: `Hello, world!` and moves to the next line.

### How to Run :
```bash
cd Questions/Begineer/0001-hello-world
go run main.go
```

