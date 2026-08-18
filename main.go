package main

import "fmt"

func main() {
	isOk := false

	if isOk {
		fmt.Println("Everything is ok")
	} else {
		fmt.Println("Not ok")
	}

	for i := 1; i <= 10; i++ {
		fmt.Println("sl", i)
	}

	// for i := range 10 {
	// 	fmt.Println("->", i)
	// }
}
