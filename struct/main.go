package main

import "fmt"

type UserData struct {
	name  string
	age   int
	class string
	job   string
}

func main() {
	var newUser UserData
	fmt.Printf("%+v\n", newUser)

	userOne := UserData{"adnan", 20, "10th", "Frontend Developer"}
	fmt.Printf("%+v\n", userOne)
	fmt.Println(userOne.age)

}
