package main

import "fmt"

func init() {
	fmt.Println("Initializing package1...")
}

func init() {
	fmt.Println("Initializing package2...")
}

func init() {
	fmt.Println("Initializing package3...")
}

func main() {
	fmt.Println("Inside the main function")
	// Init function useful to perform initialization tasks for the package before it is used.
	// Always executed before the main function, and it occurs exactly once per package
	// Even if the package imported multiple times (in multiple files)
	// It could initialize some variables, some state that is required for the package
}

// Practical Use Cases:
// - Setup Tasks
// - Configuration
// - Registering Components
// - Database Initialization (open database connection or schema migration needed for the package)
// Best Practices:
// - Avoid Side Effects
// - Initialization Order
// - Documentation
