package main

import "fmt"

func main() {

	defer fmt.Println("Hello") //defer skips the line and execute it last after all normal ones are execueted.
	defer fmt.Println("World1")
	defer fmt.Println("World2") //Reverse order of execyution

	fmt.Println("World3")
	mydefer()
}

func mydefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Print(i)
	}
}
