package basics

import "fmt"

func main() {
	// var arrayName [size]elementType

	var numbers [5]int
	fmt.Println(numbers)

	numbers[4] = 20
	fmt.Println(numbers)

	numbers[0] = 9
	fmt.Println(numbers)

	fruits := [4]string{"Apple", "Banana", "Mango", "Orange"}
	fmt.Println("Fruits array:", fruits)

	fmt.Println("Third element:", fruits[2])

	// In go, array are value types, not reference types, so when you assign an array to another array, it creates a copy of the original array.
	// Any changes made to the new array will not affect the original array.

	originalArray := [3]int{1, 2, 3}
	fmt.Println("Original array:", originalArray)
	copiedArray := originalArray

	copiedArray[0] = 10
	fmt.Println("Copied array:", copiedArray)

	for i := 0; i < len(numbers); i++ {
		fmt.Println("Element at index,", i, ":", numbers[i])
	}

	for i, v := range numbers {
		fmt.Printf("Index: %d, Value: %d\n", i, v)
	}

	// Underscore in go is used to ignore the value of a variable.
	for _, v := range numbers {
		fmt.Printf("Value: %d\n", v)
	}
	// Mapping return values to variables using underscore to ignore the first value
	_, b := someFunction()
	fmt.Println(b)

	c := 2
	_ = c // This is a way to ignore the unused variable warning in Go. The underscore is used to indicate that the variable is intentionally unused.

	// Use len to count array length
	fmt.Println("The length of numbers array is:", len(numbers))

	// Comparing Arrays
	array1 := [3]int{1, 2, 3}
	array2 := [3]int{1, 2, 3}
	// Go compare every element of the array, if all elements are equal, then the arrays are equal.

	fmt.Println("Array1 is equal to Array2:", array1 == array2)

	// Matrix
	var matrix [3][3]int = [3][3]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	fmt.Println("Matrix:", matrix)

	// Pointer and Ampersand (&) in Go
	originalArrayv2 := [3]int{1, 2, 3}
	fmt.Println("Original array:", originalArrayv2)
	var copiedArrayv2 *[3]int
	copiedArrayv2 = &originalArrayv2 // Assigning the address of originalArrayv2 to copiedArrayv2
	fmt.Println("Copied array:", copiedArrayv2)
}

func someFunction() (int, int) {
	return 1, 2
}
