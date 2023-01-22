package main

import "fmt"

func main() {
	fmt.Println("Welcome to Structs")
	// no inheritance in go lang; no super or parent
	hitesh := User{"Hitesh", "hitrsdh@go.dev", true, 21}
	fmt.Println("Hitesh:", hitesh)
	fmt.Printf("Hitesh Details are : %+v\n", hitesh)
	fmt.Printf("Hitesh Details are : %v\n Email: %v\n", hitesh.Name, hitesh.Age)

}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
