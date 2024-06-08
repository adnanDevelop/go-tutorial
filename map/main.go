package main

import "fmt"

func main() {
	fmt.Println("How to create map in golang")

	createMap := map[string]string{"name": "adnan", "age": "20", "class": "10th"}

	// delete(createMap, "name")
	for key, value := range createMap {
		fmt.Println("key:", key, "value: ", value)
	}

}
