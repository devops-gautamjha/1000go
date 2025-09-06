# 0006 - Constants and `iota` (Go-Style Enums)

---

### Problem Statement

Write a program that:

1. Declares a constant (like `Pi`)
2. Uses `iota` to implement enum-style constants such as:
   - Dev
   - Staging
   - Prod

---

### Source Code

```go
package main

import "fmt"

const (
	Dev     = iota // 0
	Staging        // 1
	Prod           // 2
)

func main() {
	const Pi = 3.14 // Constant value

	fmt.Println(Pi)
	fmt.Println(Dev)
	fmt.Println(Staging)
	fmt.Println(Prod)
}
```
### Explanation
**`const` in Go**  
`const` is used to declare values that cannot be changed after they are defined.
```go 
const Pi = 3.14
```
This is useful for values that are fixed — like configuration constants, math constants, or flags.

**What is `iota`?**
**iota** is a special keyword in Go used to create a sequence of constants that automatically increment.

```go
const (
	Dev = iota     // 0
	Staging        // 1
	Prod           // 2
)
```

This is how Go simulates enums (named constants with unique values).

- `Dev == 0`

- `Staging == 1`

- `Prod == 2`

You now have readable, consistent names for numeric constants — great for writing cleaner, safer logic.

### But Wait... What Are Enums?

In programming, enums (enumerations) are named values that represent a fixed set of options.

For example, an environment variable might only allow:

- `"dev"`

- `"staging"`

- `"prod"`

Using plain strings:

```go
env := "developement" // typo! ❌
```

That kind of mistake is hard to catch.

Instead, you define enums:

```go
const (
	Dev = iota
	Staging
	Prod
)
```

Now you use:
```go
env := Dev // ✅ much safer
```
You avoid typos like `"DEVVV"` or `"developement"` because you're picking from named constants, not user-entered strings.

### So... Can Someone Still Do This?
```go
setEnv(3) // invalid value??
```
Yes — this will compile! That’s because iota just creates integers.
If your setEnv function looks like this:
```go
func setEnv(env int) { ... }
```
Then `setEnv(Dev)` (which is `0`) and `setEnv(3)` (which is invalid) are both technically valid integers.

So how do we make it safer?


### How to Make It Safer
1. Add runtime validation
```go
func setEnv(env int) {
    if env != Dev && env != Staging && env != Prod {
        fmt.Println("Invalid environment!")
        return
    }
    // Proceed with valid env
}
```

2. Use a custom type
```go
type Environment int

const (
    Dev Environment = iota
    Staging
    Prod
)
```
Now, your function becomes:
```go
func setEnv(env Environment) { ... }
```
This improves safety and clarity — and tools like linters can help catch mistakes.


### Output
```text
3.14
0
1
2
```

### Summary
| Concept       | Explanation                                                       |
| ------------- | ----------------------------------------------------------------- |
| `const`       | Declares fixed values that never change                           |
| `iota`        | Automatically generates incrementing values inside `const` blocks |
| Enums         | Named constants that represent fixed categories or options        |
| `Dev = iota`  | Go’s way of mimicking enums                                       |
| `setEnv(Dev)` | Safe if you validate input or use custom types                    |
| `setEnv(3)`   | Still compiles, so validation is needed                           |


### Pro Tip for DevOps

Enums (and constants in general) are super helpful in DevOps tooling for:

- Environment selection (`dev`, `prod`, etc.)
- Status codes (`OK`, `FAILED`)
- Flags (`ENABLED`, `DISABLED`)
Using `iota` keeps your code cleaner, safer, and easier to maintain.

