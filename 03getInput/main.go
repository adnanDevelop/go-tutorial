package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	fmt.Println("Welcome User")

	reader := bufio.NewReader(os.Stdin) // stdin stand for standard input/ output
	fmt.Println("Rate the pizza")

	// comma ok syntax || error syntax
	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for giving rating", input)

}
