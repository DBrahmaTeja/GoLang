package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Welcome to loops")
	days := []string{"Sunday", "Tuesday", "Friday", "Wednesday", "Saturday"}
	fmt.Println(days)
	//for loop
	/*for d := 0; d < len(days); d++ { //; is mandatory
		fmt.Println(days[d])
	}*/
	//2.for loop
	/*for d := range days {
		fmt.Println(days[d])
	}*/
	/*for index, day := range days {
		fmt.Printf("index is %v and the value is %v\n", index, day)
	}*/
	rand.Seed(time.Now().UnixNano())
	rouguevalue := rand.Intn(10)
	for rouguevalue < 10 { //similar to while loop in other languages
		if rouguevalue == 2 {
			goto lco
		}
		if rouguevalue == 5 {
			break
		}
		if rouguevalue == 3 {
			rouguevalue++
			continue
		}
		fmt.Println("Value is:", rouguevalue)
		rouguevalue++
	}

lco:
	fmt.Println("Jumping at Learn Coding online")

}
