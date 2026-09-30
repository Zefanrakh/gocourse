package main

import "fmt"

func main() {

	sequence := adder()
	fmt.Println(sequence())
	fmt.Println(sequence())
	fmt.Println(sequence())
	fmt.Println(sequence())
	// The result will bew 1, 2, 3, 4 etc

	sequence2 := adder()
	fmt.Println((sequence2))
}

func adder() func() int {
	i := 0
	fmt.Println("previous value of i:", i)
	return func() int {
		i++
		fmt.Println("added 1 to i")
		return i
	}
}
