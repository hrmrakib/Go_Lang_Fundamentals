package main

import "fmt"

func main() {
	rakib := User{"Rakib", "rakib@gmail.com", true, 25}

	fmt.Println(rakib)
	fmt.Printf("In details of Rakib: %+v\n", rakib)
	fmt.Printf("Name: %v, Email: %v", rakib.Name, rakib.Email)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
