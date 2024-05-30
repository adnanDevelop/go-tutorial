package main

import (
	"fmt"
	"time"
)

func main() {
	getDate := time.Now()
	fmt.Println(getDate)
	fmt.Println("Date:", getDate.Format("01-02-2006"))
	fmt.Println("Time:", getDate.Format("15:04:05"))
	fmt.Println("Day:", getDate.Format("Monday"))

	// Create custom date
	createDate := time.Date(2024, time.May, 30, 9, 19, 0, 0, time.UTC)
	fmt.Println(createDate)

	// If we want to get the date from custom date
	fmt.Println(createDate.Format("01-02-2006"))

}
