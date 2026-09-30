package basics

import (
	"fmt"
	"os"
)

func main() {
	defer fmt.Println("Deferred statement")
	// when os.exit is called, it stops all the goroutines, and the program terminates immediately, and any deferred functions are not executed

	fmt.Println("starting the main function")

	os.Exit(1)

	// This will never be executed
	fmt.Println("End of main function")

	// Since os.Exit bypasses defrred actions to ensure that all necessary cleanup operations are performed explicitly before calling os.Exit

	// Must use sparingly, e.g. when the program must stop immediately
}
