package main

import "fmt"

func main() {
	// firstArr := [2]int{1,2}
	// secondArr := [...]string{"first", "second"}
	// secondArr[2] = "third";
	// fmt.Println(secondArr, secondArr[0])

	firstSlice := []int{1, 2, 3}
	secondSlice := []int{4, 5, 6}

	secondSlice = append(secondSlice, firstSlice...)

	fmt.Println(firstSlice, secondSlice)

}