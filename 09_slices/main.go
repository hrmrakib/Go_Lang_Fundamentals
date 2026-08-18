package main

import (
	"fmt"
	"sort"
)

func main() {
	var fruits = []string{"Apple", "Tomato", "Peach"}

	fmt.Printf("Type of the fruit: %T\n", fruits)
	fmt.Println("There is fruits", fruits)

	fruits = append(fruits, "Orange", "Mango")

	fmt.Println("New added fruits also here:", fruits)
	fmt.Println("Split fruits also here:", fruits[:1])

	highScores := make([]int, 4)

	highScores[0] = 144
	highScores[1] = 855
	highScores[2] = 365
	highScores[3] = 478
	// highScores[3] = 522

	highScores = append(highScores, 255, 369, 777)

	fmt.Println(highScores)
	sort.Ints(highScores)

	fmt.Println(highScores)

	var courses = []string{"TypeScript", "SQL", "Docker", "Nextjs", "Golang"}
	var index int = 2
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)
}
