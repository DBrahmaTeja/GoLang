package main

import (
	"fmt"
)

func main() {
	fmt.Println("Welcome to Slices:")
	/*var fruitlist = []string{"Peach"}
	fruitlist = append(fruitlist, "Apple", "Amla")
	fmt.Println("fruits:", fruitlist)
	sort.Slice(fruitlist, func(i, j int) bool {
		if fruitlist[i] < fruitlist[j] {
			return true
		} else {
			return false
		}
	})
	fmt.Println("fruits:", fruitlist)

	fruitlist = append(fruitlist[1:], fruitlist[1], "Banana")
	fmt.Println("fruits:", fruitlist)
	fmt.Println("fruits:", len(fruitlist))

	highscores := make([]int, 4)
	highscores[0] = 234
	highscores[1] = 945
	highscores[2] = 465
	highscores[3] = 867
	//highscores[4] = 987
	highscores = append(highscores, 555, 666, 331)
	fmt.Println("HighScores: ", highscores)
	sort.Ints(highscores)
	fmt.Println("highscores:", highscores)
	fmt.Println(sort.IntsAreSorted(highscores))*/

	//how to remove a value from slices based on index

	var courses = []string{"reactjs", "javascript", "swift", "python", "ruby"}
	fmt.Println("Courses: ", courses)
	index := 2
	courses = append(courses[:index], courses[index+1:]...) //index is not included
	fmt.Println("Courses: ", courses)

}
