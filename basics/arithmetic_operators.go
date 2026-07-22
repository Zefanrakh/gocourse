package basics

import (
	"fmt"
	"math"
)

func main() {
	var a, b int = 10, 3
	var result int
	result = a + b

	fmt.Println("Addition:", result)

	result = a - b
	fmt.Println("Subtraction:", result)

	result = a * b
	fmt.Println("Multiplication:", result)

	result = a / b
	fmt.Println("Division:", result)

	result = a % b
	fmt.Println("Remainder:", result)

	const pi float64 = 22.0 / 7.0
	// anything that has no decimal point, consider as an integer
	// The direction is from right to left, if we write 7, the result will be 3 not 3.0, it will be considered as an integer, and the result of 22/7 will be an integer, which is 3. Then, it will be stored to pi as an integer too (3.0, and displayed as 3)
	fmt.Println(pi)

	// Overflow with signed integers
	var maxInt int64 = 9223372036854775807
	fmt.Println(maxInt)

	maxInt = maxInt + 1
	fmt.Println((maxInt))

	// Overflow with unsigned integer, unsigned only positive
	var uMaxInt uint64 = 18446744073709551615 // max value for uint64 type
	fmt.Println(uMaxInt)

	uMaxInt = uMaxInt + 1
	fmt.Println(uMaxInt)

	// Undeflow
	var smallFloat float64 = 1.0e-323
	fmt.Println(smallFloat)

	smallFloat = smallFloat / math.MaxFloat64
	fmt.Println(smallFloat)
}
