package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Hello in Files")
	content := "This needs to go into the file."

	file, err := os.Create("./mynewfile.txt")

	checkNilError(err)
	length, err := io.WriteString(file, content)
	checkNilError(err)
	fmt.Println("Length is :", length)
	defer file.Close()
	readFile("./mynewfile.txt")
}

func readFile(filename string) {
	databyte, err := ioutil.ReadFile(filename)
	checkNilError(err)
	fmt.Println("Text Data inside the file is ", string(databyte))
}

func checkNilError(err error) {
	if err != nil {
		panic(err) //displays the error and shuts the prog
	}
}
