package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "welcome to user input"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating for the movie: ")

	//comma ok syntax or err ok syntax
	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for Rating: ", input)
	fmt.Printf("Type is %T", input)
}
