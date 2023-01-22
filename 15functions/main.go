package main

import "fmt"

func main() {
	fmt.Println("Welcome to the functions")
	greeter()
	result := adder(3, 5)
	fmt.Println("Reslut is :", result)
	res, mymsg := proAdder(2, 4, 7, 8)
	fmt.Println(mymsg, res)
}
func greeter() {
	fmt.Println("Hello from GOLANG!!")
}
func adder(a int, b int) int {
	return a + b
}
func proAdder(values ...int) (int, string) {
	total := 0
	for _, val := range values {
		total += val
	}
	return total, "Hi Pro result:"
}
