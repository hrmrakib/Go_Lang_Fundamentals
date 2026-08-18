package main

import "fmt"

func main() {
	age := 511

	if age > 150 {
		fmt.Println("Sorry, you are NOT exist in the world")
	} else if age > 18 {
		fmt.Println("Hey, you are adult")
	} else {
		fmt.Println("You are NOT adult")
	}

	if n := 6; n > 5 {
		fmt.Println("Mature number")
	} else {
		fmt.Println("Immature number")
	}

}
