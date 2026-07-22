package main

import "fmt"

// package name short, concise, and in lowercase

type EmployeeGoogle struct {
	FirstName string
	LastName  string
	Age       int
}

type EmployeeApple struct {
	FirstName string
	LastName  string
	Age       int
}

func main() {
	// PascalCase
	// Eg. CalculateArea, UserInfo, NewHTTPRequest
	// Structs, interfaces, enums

	// snake_case
	// Eg user_id, fist_name, http_request

	// UPPERCASE
	// Use case is Constants
	const MAXRETRIES = 5

	// mixedCase
	// Eg. javaScript, htmlDocument, isValid
	var employeeID = 1001
	fmt.Println("EmployeeID: ", employeeID)
}
