package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Welcome to switch case in go lang:")
	rand.Seed(time.Now().UnixNano())
	dicenumber := rand.Intn(6) + 1
	fmt.Println("Value of Dice is:", dicenumber)
	switch dicenumber {
	case 1:
		fmt.Println("Value is 1 and You can Open")

	case 2:
		fmt.Println("Value is 2 and You can Open")

	case 3:
		fmt.Println("Value is 3 and You can Open")
		fallthrough
	case 4:
		fmt.Println("Value is 4 and You can Open")
		fallthrough //continues the next case also
	case 5:
		fmt.Println("Value is 5 and You can Open")

	case 6:
		fmt.Println("Value is 6 and You can Open")
	}

}
