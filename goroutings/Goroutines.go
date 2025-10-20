package main

import (
	"fmt"
	"time"
)

func display(msg string) {
	for i := 1; i <= 5; i++ {
		fmt.Println(msg, i)
	}
} 

func main() {
	go display("Hello, Goroutine!")
	time.Sleep(1000)
    display("Hello, Main!")
    display("Hello, Main!")
    display("Hello, Main!")
    display("Hello, Main!")
    display("Hello, Main!")
}