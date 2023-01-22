package main

import "fmt"

func main() {
	fmt.Println("Welcome to Structs")
	// no inheritance in go lang; no super or parent
	hitesh := User{"Hitesh", "hitrsdh@go.dev", true, 21}
	// fmt.Println("Hitesh:", hitesh)
	// fmt.Printf("Hitesh Details are : %+v\n", hitesh)
	// fmt.Printf("Hitesh Details are : %v\n Email: %v\n", hitesh.Name, hitesh.Age)
	hitesh.GetStatus()
	hitesh.NewMail()
	fmt.Println("Email of this user is: ", hitesh.Email)

}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus() {
	fmt.Println("Status:", u.Status)
}
func (u User) NewMail() {
	u.Email = "test@go.dev"
	fmt.Println("Email of this user is: ", u.Email)
}
