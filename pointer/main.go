package main

import "fmt"

func main() {

	// var myNumber *int First example without assigning a value
	assignValue := 23

	getValue := &assignValue // Here i'm creating a pointer which also referencing some memory

	// fmt.Println(myNumber, assignValue, getValue, *getValue)
	fmt.Println(*getValue) // ye humain jo getValue k ander memory h uske ander jo value hogi we return kery ga
	fmt.Println(getValue)  // ye humain memory address return kery ga

}
