package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to Time Study!")
	presentTime := time.Now()
	fmt.Println(presentTime)

	fmt.Println(presentTime.Format("01-02-2006 "))
	fmt.Println(presentTime.Format("Monday"))
	fmt.Println(presentTime.Format("15:04:05"))
	fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday"))

	createdDate := time.Date(2021, time.August, 17, 13, 13, 0, 0, time.UTC)
	fmt.Println(createdDate)
	fmt.Println(createdDate.Format("01-02-2006"))
	fmt.Println(createdDate.Format("Monday"))
	fmt.Println(createdDate.Format("01-02-2006 Monday 15:04:05"))

}
