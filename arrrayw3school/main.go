package main

import "fmt"

// func main() {
// 	var students = [2]string{"Adnan", "Hanan"}
// 	walrusStudends := [2]int{1, 2}
// 	 fmt.Println(students, walrusStudends)

// 	//  Access elements of an array
// 	fmt.Println(students[0], walrusStudends[1])

// 	// Changes second element value
// 	walrusStudends[1] = 50
// 	fmt.Println(walrusStudends)

// 	/* Initiliaze only the second and third elements of the array*/
// 	addNewElements := [5]int{1: 10, 4: 20}
// 	addNewElements[3] = 30 	// adding new value

// 	fmt.Println(addNewElements)
// 	fmt.Println(addNewElements[1]) // get specific value using index number
// 	fmt.Println(len(addNewElements)) //Checking the length of an array

// }

func main() {
	var firstStrArr = [3]string{"first", "second", ""}
	var firstInArr = [3]int{1, 2, 3}
	firstStrArr[2] = "third";
	fmt.Println(firstStrArr, "\n", firstInArr)
}
