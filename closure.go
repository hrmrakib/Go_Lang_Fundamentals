package main
import "fmt"

func multiplier(fector int) func(int) int {
	fmt.Println("multiplier called", fector)
	return func(x int) int {
		fmt.Println("anonymous function called", x)
		return x * fector
	}
}

func main() {
	double := multiplier(2)
	// triple := multiplier(5)
	// four := multiplier(10)

	fmt.Println(double(5))
	// fmt.Println(triple(5))
	// fmt.Println(four(10))
}