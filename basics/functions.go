package basics

import "fmt"

func main() {
	// func<name>(parameters list) returnType {
	// code block
	// return value
	// }

	// fmt.Println(), 'P' is uppercase, because it's public function

	// If return statement is omitted, functions return default value for the return type.

	sum := add(1, 2)
	fmt.Println(sum)

	// argument that passes to the function are COPIED into the function's parameters

	greet :=
		func() {
			fmt.Println("hello anonymous function")

		}
	greet()

	result := applyOperation(5, 3, add)
	fmt.Println(result)

	multiplyBy2 := createMultiplier(2)
	fmt.Println("6 * 2 = ", multiplyBy2(6))
}

func add(a, b int) int {
	return a + b
}

// Function that takes a function as an argument
func applyOperation(x int, y int, operation func(int, int) int) int {
	return operation(x, y)
}

// Function that returns a function
func createMultiplier(factor int) func(int) int {
	return func(x int) int {
		return x * factor
	}
}
