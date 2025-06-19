package main

import "fmt"

func main() {
	var number int
	
	fmt.Print("Enter a number: ")
	fmt.Scanln(&number)
	if number > 1 && number < 10  {
		fmt.Print("True")
	} else {
		fmt.Print("False")
	}

}