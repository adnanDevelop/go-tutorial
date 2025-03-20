package main

import "fmt"

func main() {
	testingOne := "One"
	testingTwo := 2
	testingThree := false

	/*
		There are three ways to display output
		1: fmt.Println()
		2: fmt.Printf()
		3: fmt.Fprintf()
	*/

	/*
	 	fmt.Print("hello world")
	 	Ager 2 variables ko at a time print kerain toh ye dono aik hi line mein show kerta ha aur dono mein koi space nai add kerta humain /n iska use kerna parhta h ager new line create kerni ho toh ager space add kerni ho toh  " " empty string ka use kerty h. Ye tb space add kerta h jb string k ilawa koi aur type hoti h
	 */
	fmt.Print("Without space and new line ", testingOne, testingTwo, testingThree, "\n")
	fmt.Print("With space ", testingOne, " ", testingTwo, " ", testingThree, "\n")
	fmt.Print("With new line \n", testingOne, "\n", testingTwo, "\n", testingThree, "\n")

	/*
		fmt.Println("hello world")
		Ye bhi Print ki tarah hi kaam kerta h lekin by default space aur new line add kerta h
	
	*/

	fmt.Println("With spaces and new line", testingOne, testingTwo, testingThree)


	/*
		fmt.Printf("hello world")
		Ye formated output dene k liye use kerty h like %v %s %f %t is tarah k aur ye new line nai create kerta khud add kerna perhta h "\n"
	
	*/
	fmt.Printf("Without space and new line %v %v %v \n", testingOne, testingTwo, testingThree)
	fmt.Printf("the type of variable: %T and value is: %v", testingOne, testingOne)
	fmt.Printf("the type of variable: %T and value is: %v", testingTwo, testingTwo)
	fmt.Printf("the type of variable: %T and value is: %v", testingThree, testingThree)

	

}