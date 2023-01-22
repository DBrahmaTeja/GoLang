package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome to our pizza app")
	fmt.Println("Rate the pizza")
	reader := bufio.NewReader(os.Stdin)

	input, _ := reader.ReadString('\n')
	fmt.Println("Rating ,", input)
	numRating, err := strconv.ParseInt(strings.TrimSpace(input), 10, 32)
	if err != nil {
		fmt.Println(err)
	} else {

		fmt.Println("Num Rating: ", numRating+1)
	}
}
