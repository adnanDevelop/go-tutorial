package main

import "fmt"

func main() {
	for i := 0; i < 5; i++ {
		fmt.Println("the value is: ", i)
	}

	// loop on an array using range
	fruits := []string{"apple", "banana", "orange", "mango"}
	for _, fruits := range fruits {
		fmt.Println( fruits)
	}

}
