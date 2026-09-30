package basics

import "fmt"

func main() {

	// Recover only useful inside defer functions

	// Used to manage behavior of a panicking go routine to avoid abrupt termination.
	process()
	fmt.Println("Returned from process")
}

func process() {
	defer func() {
		if r := recover(); r != nil {
			// r := recover()
			// if r != nil {
			fmt.Println("Recovered from panic:", r) // usually this thing write to error handling
		}
	}()

	fmt.Println("Start Process")
	panic("Something went wrong")
	fmt.Println("End Process")

}

// Practical use cases:
// - Graceful Recovery
// - Cleanup
// - Logging and Reporting

// Panics and recover should be used sparingly, and only for exceptional unrecoverable errors.
// Do not overuse, as it can make your code harder to reason about.
