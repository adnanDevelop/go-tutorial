package main

import "fmt"

func main() {

	// var myNumber *int First example without assigning a value
	assignValue := 23

	getValue := &assignValue

	// fmt.Println(myNumber, assignValue, getValue, *getValue)
	fmt.Println(*getValue)

}
