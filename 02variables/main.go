package main

import "fmt"

//jwtToken := 1250;	    this gives an error because we can't decalre variable like this globally it work's in any method or in function

const MyName = "Adnan Tariq" //If we are creating a variable with constant then we have to put the first letter of constant variable to capital

// := we can call this volrus operator

func main() {
	var name string = "adnan" //STRING DATA
	fmt.Println(name)
	fmt.Printf("The variable of type %T \n", name)

	var isLogIn bool = false //BOOLEAN DATA
	fmt.Println(isLogIn)
	fmt.Printf("The variable of type %T  \n", isLogIn)

	var numbers int = 1200 //BOOLEAN DATA
	fmt.Println(numbers)
	fmt.Printf("The variable of type %T \n", numbers)

	var floatNumbers = 12.50 //BOOLEAN DATA
	fmt.Println(floatNumbers)
	fmt.Printf("The variable of type %T \n", floatNumbers)

	// WITHOUT TYPE
	const withoutType = "this is without type"
	fmt.Println(withoutType)

	//CREATE VARIABLE WITHOUT VAR KEWORD

	withouVarKeyword := "this variable is declared without var keyword"
	fmt.Println(withouVarKeyword)

}
