package main

import "fmt"

func main() {
	roll := map[int]string{
		1:  "Aman",
		2:  "Arjun",
		3:  "Bhaskar",
		4:  "Chandu",
		5:  "Divakar",
		6:  "Gautam",
		7:  "Hardik",
		8:  "Ishita",
		9:  "Julie",
		10: "Krishna",
	}

	// Printing the Map

	for rollnum, name := range roll {
		fmt.Printf("The Student of rollno. %d is %s \n", rollnum, name)
	}

	delete(roll, 9)

	fmt.Println()

	for rollnum, name := range roll {
		fmt.Printf("The Student of rollno. %d is %s \n", rollnum, name)
	}

}
