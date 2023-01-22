package main

import "fmt"

func main() {
	fmt.Println("Welcome to Class on pointers")
	cons := 2
	var ptr = &cons
	fmt.Println("Address :", ptr)
	fmt.Println("Value is:", *ptr)
	*ptr = *ptr + 10
	fmt.Println("Address :", ptr)
	fmt.Println("Value is:", *ptr)
}
