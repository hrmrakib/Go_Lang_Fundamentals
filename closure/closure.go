package main

import "fmt"

func add(a, b int) int {
	return a + b
}

func sum() {
	var a, b, c int = 10, 20, 30
	sum := a + b + c
	fmt.Println(sum)
}

func info() {
	var a, b = 10, "hello"
	c, d := 20, "world"

	fmt.Println(a, b, c, d)
}

func myself() {
	var (
		name     string = "Rakib 2.0"
		age      int    = 24
		isMarrid bool   = false
	)

	fmt.Println(name, age, isMarrid)
}

func main() {
	// var x int = add(5, 6)
	// fmt.Println(x)

	// sum()
	// info()
	myself()
}
