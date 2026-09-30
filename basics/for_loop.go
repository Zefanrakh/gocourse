package basics

import "fmt"

func main() {
	// Simple iteration
	// for i := 1; i <= 5; i++ {
	// 	fmt.Println(i)
	// }

	// Collection iteration
	// numbers := []int{1, 2, 3, 4, 5, 6} // This is slices
	// for index, value := range numbers {
	// 	// Printf if use placeholder
	// 	// %v for general value
	// 	// %d for number
	// 	fmt.Printf("Index: %d, Value: %d\n", index, value)
	// }

	// Continue & Break
	// for i := 1; i <= 10; i++ {
	// 	if i%2 == 0 {
	// 		continue
	// 	}
	// 	fmt.Println("Odd Number:", i)
	// 	if i == 5 {
	// 		break
	// 	}
	// }

	// Outer & Inner
	// rows := 5
	// for i := 1; i <= rows; i++ {

	// 	for j := 1; j <= rows-i; j++ {
	// 		fmt.Print(" ")
	// 	}
	// 	for k := 1; k <= 2*i-1; k++ {
	// 		fmt.Print("*")
	// 	}
	// 	fmt.Println()

	// }

	// Compatibility to range over integers directly in for loops in go 1.22
	for i := range 10 {
		fmt.Println(10 - i)
	}
}
