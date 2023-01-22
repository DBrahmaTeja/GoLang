package main

import "fmt"

func main() {
	fmt.Println("Welcome to control flow: IF_ELSE:")
	loginCount := 2
	var result string
	if loginCount < 10 {
		result = "Regular User"
	} else if loginCount > 10 {
		result = "Something Else"
	} else {
		result = "Special User"
	}
	fmt.Println("Result is :", result)

	if 9%2 == 0 {
		fmt.Println("Even Number")
	} else {
		fmt.Println("Odd Number")
	}

	if num := 3; num < 10 { //local variable is created not globally
		fmt.Println("Single Digit")
	} else {
		fmt.Println("Double Digit")
	}
	//fmt.Println("NUmber is :", num)

}
