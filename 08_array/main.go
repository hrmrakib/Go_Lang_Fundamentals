package main

import "fmt"

func main() {
	var lang [4]string

	lang[0] = "C"
	lang[1] = "C++"
	lang[3] = "C#"

	fmt.Println("Array's value:", lang)

	var letters = []string{"A", "I", "N", "R", "S"}

	fmt.Println("Letters here", letters)
}
