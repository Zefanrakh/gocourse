package basics

import (
	"errors"
	"fmt"
)

func main() {
	// func functionName(parameter1 type1, parameter2 type2, ...) (
	// 	returnType1, returnType2, ...) {
	// 		return returnvalue1, returnValue2, ...
	// 	}

	q, r := divide(10, 3)
	fmt.Printf("Quotient: %d. Remainder: %d\n", q, r)

	result, err := compare(3, 3)
	if err != nil {
		fmt.Println("Error: ", err)
	} else {
		fmt.Println((result))
	}
}
func divide(a, b int) (int, int) {
	quotient := a / b
	remainder := a % b
	return quotient, remainder
}

func compare(a, b int) (string, error) {
	if a > b {
		return "a is greater than b", nil
	} else if b > a {
		return "b is greater than a", nil
	} else {
		return "", errors.New("Unable to compare which is greater")
	}
}

// Named return value
func divideNamed(a, b int) (quotient int, remainder int) {
	quotient = a / b
	remainder = a % b
	// Because go compiler smart enough to detect that we have already declared name return value in the return parentheses
	return
}
