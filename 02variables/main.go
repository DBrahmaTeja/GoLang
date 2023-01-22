package main

import "fmt"

const Pi float64 = 3.14167894543

func main() {
	var username string = "Brahma Teja"
	fmt.Println(username)
	fmt.Printf("Variable is of type: %T\n", username)

	var isLoggedIn bool = false
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type: %T\n", isLoggedIn)
	//no var style can't be used outside of the function
	numberofuser := 100.0
	fmt.Println("NO.of Users:", numberofuser)
	fmt.Printf("Variable is of type: %T\n", numberofuser)

	fmt.Println("Pi:", Pi)
}
