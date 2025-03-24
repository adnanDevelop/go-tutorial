package main

import "fmt"

func testingFunc(firstName, lastName string) {
	fmt.Println("Hello", firstName, lastName)
}

func main() {
	testingFunc("Adnan", "Tariq")
}
