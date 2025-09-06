# 1000Go

A collection of **1000 Go language exercises** designed with a focus on **DevOps engineers** and **Go beginners**.

Each exercise is solved step-by-step, with:

- ✅ Clean Go code (`main.go`)
- 📄 Descriptive explanations (`README.md`)
- 🖨️ Expected outputs (`output.txt`)

---

## 📁 Folder Structure

Each question lives in its own subfolder:

```text
Questions
    └── Begineer
        ├── 0001-hello-world
        │   ├── main.go
        │   ├── output.txt
        │   └── README.md
        ├── 0002-cmd-line-arguments
        │   ├── main.go
        │   ├── output.txt
        │   └── REAMDME.md
        ├── 0003-os.args
        ├── ...
```


- `main.go`: the Go program
- `output.txt`: sample output when the program runs
- `README.md`: explanation of the problem and how it works

---

## 🚀 Getting Started

### 🔧 Requirements

Make sure you have Go installed.  
If not, install it from: https://golang.org/dl/

Verify your Go installation:

```bash
go version
```

### ▶️ Running Any Program

Navigate to any question directory and run:
```bash 
go run main.go
```
**For example**
```bash
cd Questions/Begineer/0001-hello-world
go run main.go
```

## 💡 Who This Repo Is For

- Developers and DevOps engineers getting started with Go
- Anyone who prefers learning by doing
- People who want clean, explained examples

## 📚 Topics Covered So Far
1. [Go syntax and program structure](./Questions/Begineer/0001-hello-world)
2. [Command-line arguments](./Questions/Begineer/0002-cmd-line-arguments)
3. [Working with `os.Args`](./Questions/Begineer/0003-os.args)
4. [Print functions: `fmt.Println`, `fmt.Printf`, `fmt.Sprintf`](./Questions/Begineer/0004-demonstrate-print-statements)
5. [Variables, constants, and types](./Questions/Begineer/0005-Variables-and-types)
6. [Constants and `iota` enums](./Questions/Begineer/0006-Constants-and-iota)
7. [Arrays](./Questions/Begineer/0007-print-array)
8. [Slices and their behavior](./Questions/Begineer/0008-slices-and-functions)
9. [Looping with `range`](./Questions/Begineer/0009-for-and-range)
10. [Working with maps and safe lookups](./Questions/Begineer/0010-check-for-missing-key-in-map)

_More coming soon..._

## 🧭 Naming Convention

Each folder follows this format:
```text
000X-title-of-the-question
```
This helps keep questions sorted and readable.

# 👨‍💻 Author

Created by [Gautam Jha]
Feel free to star 🌟 the repo or suggest improvements!