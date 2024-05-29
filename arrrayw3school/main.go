package main

import "fmt"

func main() {
	/*
		var students = [2]string{"Adnan", "Hanan"}
		walrusStudends := [2]int{1, 2}
		 fmt.Println(students, walrusStudends)


		Access elements of an array
		fmt.Println(students[0], walrusStudends[1])


		Changes second element value
		walrusStudends[1] = 50
		fmt.Println(walrusStudends)

		Initiliaze only the second and third elements of the array
	*/

	addNewElements := [5]int{1: 10, 4: 20}
	fmt.Println(addNewElements)
	fmt.Println(len(addNewElements)) //Checking the length of an array

}
