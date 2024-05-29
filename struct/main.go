package main

import "fmt"

type UserData struct {
	name  string
	age   int
	class string
	job   string
}

func main() {
	var userOne UserData

	userOne.name = "Adnan"
	userOne.age = 21
	userOne.class = "12th"
	userOne.job = "Frontend Web Developer"
	fmt.Println("Name:", userOne.name)
	fmt.Println("Age:", userOne.age)
	fmt.Println("Class:", userOne.class)
	fmt.Println("Job:", userOne.job)

}
