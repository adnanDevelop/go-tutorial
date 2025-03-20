package main

import "fmt"

var globalVariable string = "This is global variable"
var testing string 
func main () {
	fmt.Println(testing)
	// Variables with types
	var numberVar int = 20;
	var isAuth bool = true;
	var myName  = "Variable in the function"
	var assignLater string
	
	fmt.Println(globalVariable, "\n", myName,  )
	fmt.Println("Number type variable:", numberVar, "\n", "Boolean type variable:", isAuth)
	testing  = "Declaraing outside of the function and assigning value inside the function"



	// Variables without assigning values
	var noName string;
	var noNumber int;
	var noBool bool;
	assignLater = "This is assigned later"

	fmt.Println( "Variable without assigning value", noName, noNumber, noBool)
	fmt.Println( "This is assign later", assignLater)



	// Variables using short form variable declaration We can't use this outside of the function. It's not possible to declare a variable using := without assigning value to it
	lastName := "Tariq";
	lastNum := 20;
	notAuth := false;

	fmt.Println("Variable using short form variable declaration", lastName, lastNum, notAuth)

}