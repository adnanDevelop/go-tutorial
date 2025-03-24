package main

import (
	"fmt"
	"sort"
)

func main() {

	// Creating slice using comon method
	var mySlice = []string{"Adnan", "hanan"}
	fmt.Println(mySlice, len(mySlice), cap(mySlice)) // len return the length of slice while cap return the capacity of slice

	var cutSlice = append(mySlice, "Hamad", "Wajid")
	fmt.Println(cutSlice[1:3], "cutSlice")

	// Create a slice from array
	var newArr = [4]int{11, 22, 33, 44}
	var sliceFromNewArr = newArr[1:3]
	fmt.Println(sliceFromNewArr)

	// Create slice using make() function
	makeSlice := make([]int, 1, 2)
	fmt.Println(makeSlice)


	var sortNumber = []int{2342, 213, 345, 100, 50}

	sort.Ints(sortNumber)
	fmt.Println(sortNumber, "sorted Number")
	fmt.Println(sort.IntsAreSorted(sortNumber), "sorted slice")

	// How to delete element from slice
	var deleteElement = []string{"html", "Css", "Javascript", "Bootstrap", "React"}
	index := 2

	deleteElement = append(deleteElement[:index], deleteElement[index+1:]...)
	fmt.Println(deleteElement)
}
