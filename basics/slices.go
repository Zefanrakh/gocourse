package basics

import (
	"fmt"
	"slices"
)

func main() {
	// slices dont have length defined at declaration
	// var slicesName []elementType
	// var numbers []int
	// var numbers1 = []int{1, 2, 3}

	// numbers2 := []int{9, 8, 7}

	// slice := make([]int, 5)

	a := [5]int{1, 2, 3, 4, 5}
	sliceArray := a[1:4] // create slice from array, start at index 1, but end before index 4

	fmt.Println("Slice from array:", sliceArray)

	sliceArray = append(sliceArray, 6) // append to slice, this will create a new underlying array if the capacity is exceeded
	fmt.Println("Slice after append:", sliceArray)

	sliceCopy := make([]int, len(sliceArray))
	fmt.Println("Slice copy:", sliceCopy)

	copy(sliceCopy, sliceArray) // copy elements from sliceArray to sliceCopy
	fmt.Println("Slice copy after copy:", sliceCopy)

	// var nilSlice []int // nil slice

	for i, v := range sliceArray {
		fmt.Printf("Index: %d, Value: %d\n", i, v)
	}

	if slices.Equal(sliceArray, sliceCopy) { // use to compare slices, cannot use '==' like arrays, because slices are reference types
		fmt.Println("Slices are equal")
	}

	// Multidimensional slices
	twoD := make([][]int, 3) // the length is always for the outer slice
	fmt.Println("Two-dimensional slice before initialization:", twoD)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("Two-dimensional slice:", twoD)

	//  slice[low:high] // slice the slice from index low to high-1
	sliceArray2 := sliceArray[1:3] // create a new slice from sliceArray, start at index 1, but end before index 3
	fmt.Println("Slice from slice:", sliceArray2)

	fmt.Println("The capacity of sliceArray2 is:", cap(sliceArray2)) // use cap to get the capacity of the slice
	fmt.Println("The capacity of sliceArray is:", cap(sliceArray))   // use cap to get the capacity of the slice

	// When make slice, the third parameter is the capacity of the slice. We can use cap to prevent reallocation of the underlying array when appending to the slice.
}
