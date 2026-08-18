package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to time study of Golang!")

	presentTime := time.Now()

	fmt.Println("time", presentTime)
	fmt.Println("time is", presentTime.Format("01-02-2002"))
}
