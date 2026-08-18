package main

import "fmt"

func main() {

	// var ptr *int
	// fmt.Println("Value of the pointer", ptr) // <nil>

	n := 25
	var prt = &n

	fmt.Println("Address of the pointer", prt)
	fmt.Println("Value of the pointer", *prt)

	*prt = *prt * 2

	fmt.Println("Value of the pointer", *prt)
}
