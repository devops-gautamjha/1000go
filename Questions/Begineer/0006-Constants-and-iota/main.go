package main

import "fmt"

const (
	Dev     = iota // 0
	Staging        // 1
	Prod           // 2
)

func main() {
	const Pi = 3.14 // constant using

	fmt.Println(Pi)
	fmt.Println(Dev)
	fmt.Println(Staging)
	fmt.Println(Prod)
}
