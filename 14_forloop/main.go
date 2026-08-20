package main

import "fmt"

func main() {
	gal := []string{"Nowshin", "Mim", "Sadiya", "Jaima"}

	// for i := 0; i < len(gal); i++ {
	// 	fmt.Println(gal[i])
	// }

	// for g := range gal {
	// 	fmt.Println(gal[g])
	// }
	for _, g := range gal {
		fmt.Println(g)
	}

	rougueValue := 1

	for rougueValue <= 10 {
		if rougueValue == 2 {
			goto lco
		}

		if rougueValue == 5 {
			rougueValue++
			continue
		}

		fmt.Println(rougueValue)
		rougueValue++
	}

lco:
	fmt.Println("I am here!")
}
