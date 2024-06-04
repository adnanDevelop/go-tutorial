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
	fmt.Println(cutSlice[1:3])

	// Create a slice from array
	var newArr = [4]int{11, 22, 33, 44}
	var sliceFromNewArr = newArr[1:3]
	fmt.Println(sliceFromNewArr)

	// Create slice using make() function
	makeSlice := make([]int, 1, 2)
	fmt.Println(makeSlice)

	days := "thursday"

	switch days {
	case "monday":
		fmt.Println(false)
	case "tuesday":
		fmt.Println(false)
	case "wednesday":
		fmt.Println(false)
	case "thursday":
		fmt.Println(true)
	default:
		fmt.Println("there is no match")
	}

	/*
		newNumber := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}
		for i := 0; i < len(newNumber); i++ {
			fmt.Println(8*i, "=", 8*i)
		}
	*/

	fruit := []string{"Apple", "Banana", "Mango", "Orange", "Grapes"}

	for _, value := range fruit {
		if value == "Banana" {
			continue
		} else if value == "Orange" {
			break
		}

		fmt.Println(value)
	}

	var sortNumber = []int{2342, 213, 345, 100, 50}

	sort.Ints(sortNumber)
	fmt.Println(sortNumber)
	fmt.Println(sort.IntsAreSorted(sortNumber))

}
