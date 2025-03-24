package main

import "fmt"

func main() {
	day := "Monday"

	// Testing condition using if else
	if day == "Monday" {
		fmt.Println("Today is Monday")
	} else if day == "Tuesday" {
		fmt.Println("Today is Tuesday")
	} else if day == "Wednesday" {
		fmt.Println("Today is Wednesday")
	} else if day == "Thursday" {
		fmt.Println("Today is Thursday")
	} else if day == "Friday" {
		fmt.Println("Today is Friday")
	} else if day == "Saturday" {
		fmt.Println("Today is Saturday")
	} else if day == "Sunday" {
		fmt.Println("Today is Sunday")
	} else {
		fmt.Println("Invalid day")
	}

	// Using switch case
	switch day {
	case "Monday":
		fmt.Println("Today is Monday")
	case "Tuesday":
		fmt.Println("Today is Tuesday")
	case "Wednesday":
		fmt.Println("Today is Wednesday")
	case "Thursday":
		fmt.Println("Today is Thursday")
	case "Friday":
		fmt.Println("Today is Friday")
	case "Saturday":
		fmt.Println("Today is Saturday")
	case "Sunday":
		fmt.Println("Today is Sunday")
	default:
		fmt.Println("Invalid day")
	}

	newDay := 5
	switch newDay {
	case 1, 3, 5:
		fmt.Println("Odd day")
	case 2, 4, 6:
		fmt.Println("Even day")
	default:
		fmt.Println("Invalid day")
	}

}
