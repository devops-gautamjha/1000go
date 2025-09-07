# delete-map-keys

### 📝 Problem Statement

**Delete keys from a map and show the effect on iteration.**

In Go, we are asked to:

* Create a map with several key-value pairs.
* Iterate over the map and print the contents.
* Delete a specific key.
* Iterate again and show how the map has changed after deletion.

---

### Explanation

This program demonstrates how deleting a key from a map in Go affects its content and iteration.

Steps:

1. A `map[int]string` is created, representing student roll numbers and names.
2. The map is printed in full using a `for range` loop.
3. A specific key (`9`, representing "Julie") is deleted using the `delete()` function.
4. The map is printed again to show that the key-value pair for roll number 9 has been removed.

---

###  Source Code

```go
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

	// Delete one entry
	delete(roll, 9)

	fmt.Println()

	// Printing the Map after deletion
	for rollnum, name := range roll {
		fmt.Printf("The Student of rollno. %d is %s \n", rollnum, name)
	}
}
```

---

### Concept: Order of Maps in Go

Maps in Go are **unordered** collections. This means:

* The order in which elements appear during iteration is **not guaranteed** and can vary between runs.
* Deleting a key does not shift elements or preserve any original order.
* Go intentionally randomizes map iteration order to prevent developers from relying on it.

Thus, you should **not assume** that the map will print keys in the order they were inserted or in numeric order.

---

### Output

```
The Student of rollno. 1 is Aman 
The Student of rollno. 2 is Arjun 
The Student of rollno. 3 is Bhaskar 
The Student of rollno. 4 is Chandu 
The Student of rollno. 5 is Divakar 
The Student of rollno. 6 is Gautam 
The Student of rollno. 7 is Hardik 
The Student of rollno. 8 is Ishita 
The Student of rollno. 9 is Julie 
The Student of rollno. 10 is Krishna 

The Student of rollno. 1 is Aman 
The Student of rollno. 2 is Arjun 
The Student of rollno. 3 is Bhaskar 
The Student of rollno. 4 is Chandu 
The Student of rollno. 5 is Divakar 
The Student of rollno. 6 is Gautam 
The Student of rollno. 7 is Hardik 
The Student of rollno. 8 is Ishita 
The Student of rollno. 10 is Krishna 
```

> Note: The actual **order** may vary each time you run the program due to the unordered nature of maps in Go.

---

Let me know if you want the same structure for other problems too.
