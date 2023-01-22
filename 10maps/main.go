package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("WELCOME TO MAPS IN GOLANG")
	languages := make(map[string]string) //map[key type]value type
	languages["js"] = "javascript"
	languages["ru"] = "ruby"
	languages["py"] = "python"
	fmt.Println("Languages\n", languages)
	fmt.Println("JS shorts for ", languages["js"])
	delete(languages, "ru")
	fmt.Println("Languages\n", languages)
	//loops
	for key, value := range languages {
		fmt.Printf("For key %v, Value is %v\n", strings.ToUpper(key), value)
	}
}
