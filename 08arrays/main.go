package main

import "fmt"

func main() {
	fmt.Println("Welcome to Arrays")

	var fruitlist [4]string
	fruitlist[0] = "Apple"
	fruitlist[1] = "Mango"
	fruitlist[3] = "Peach"
	fmt.Println("Fruits:", fruitlist)
	fmt.Println("Fruits:", len(fruitlist))

	var vegList = [3]string{"Potato", "", "Tomato"}
	fmt.Println("Veggies: ", vegList)
}
