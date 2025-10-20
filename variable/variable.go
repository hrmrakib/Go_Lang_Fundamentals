package main

import "fmt"

func variableScopeDemo() {
	// Local variable declaration and initialization with same type
	var a, b, c int = 1, 2, 3
	fmt.Println("a:", a, "b:", b, "c:", c)
	fmt.Println("---------------------")

	// Local variable declaration and initialization with different types
	var x, y = 10, "hello"
	m, n := 3.14, true

	fmt.Println("x:", x, "y:", y)
	fmt.Println("m:", m, "n:", n)
	fmt.Println("---------------------")

	// Grouped variable declaration with mixed types
	var (
     d int
     e int = 1
     f string = "hello"
   )
 
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
}

func main() {
	variableScopeDemo()
}