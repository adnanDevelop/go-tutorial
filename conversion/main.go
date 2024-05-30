package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	getInput := bufio.NewReader(os.Stdin)
	fmt.Println("Please rate the pizza between 1 to 5")

	response, _ := getInput.ReadString('\n')
	fmt.Println(`Thanks for giving rate`, response)

	getAnswer, error := strconv.ParseFloat(strings.TrimSpace(response), 64)

	if error != nil {
		fmt.Println(error)
	} else {
		fmt.Println("Added 1 number to it: ", getAnswer+1)
	}

}
