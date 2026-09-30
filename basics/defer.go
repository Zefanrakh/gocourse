package basics

import "fmt"

func main() {

	// The concept similar to finally

	// Go routines are functions which run in the background, which are
	// running concurrently in the background, and they are not part of
	// the main thread

	// So any function which is a go routine is thrown to the back so
	// that it finishes off its work, not in the main thread, not blocking
	// the main thread but in the background, and then comes back and
	// joins the main thread, once it's finished.

	// The function call is evaluated immediately, but the execution is
	// deferred until the surrounding function returns.

	// So we can make a defer function at the beginning, but it will only
	// execute once the surrounding function returns.

	// So the defer function is always part of another function.

	// The surrounding function means the function the encloses the defer function.

	process(10)
}

func process(i int) {
	// Arguments of the deferred function are evaluated immediately
	defer fmt.Println("Deffered i value:", i)
	// Print in reverse order (LIFO - Last In, First Out)
	defer fmt.Println("First deferred statement executed")
	defer fmt.Println("Second deferred statement executed")
	defer fmt.Println("Third deferred statement executed")
	fmt.Println("Normal execution statement")
	i++
	fmt.Println("Value of i:", i)
}

// Practical use cases:
// - resource cleanup
// - unlocking mutexes
// - logging and tracing
// Best Practices:
// - Keep Deferred Actions Short
// - Understang Evalutaion Timing
// - Avoid Complex Control Flow
