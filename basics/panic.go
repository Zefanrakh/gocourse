package basics

import "fmt"

func main() {
	// panic(interface{}) it can input any type

	process(10)

	process(-3)
}

func process(input int) {

	defer fmt.Println("Deferred 1")
	defer fmt.Println("Deferred 2")
	if input < 0 {
		fmt.Println("Before Panic")
		panic("Input must be a non-negative number")
		// fmt.Println("After Panic")
		// And then it gave us error stack

		// defer fmt.Println("Deferred 3")
	}
	fmt.Println("Processing input:", input)
}
