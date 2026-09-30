package basics

import "fmt"

func main() {
	// switch case in go is switch case default
	// switch case in other languages is switch case break default
	// switch expression {
	// case value1:
	// case value2:
	// default:
	// }

	// To use fallthrough in go, you need to use the fallthrough keyword. It will execute the next case even if the condition is not met.

	fruit := "apple"

	switch fruit {
	case "apple":
		println("This is an apple.")
	case "banana":
		println("This is a banana.")
	default:
		println("Unknown fruit.")
	}

	// Multiple conditions
	day := "Monday"
	switch day {
	case "Monday", "Tuesday", "Wednesday", "Thursday", "Friday":
		println("It's a weekday.")
	case "Saturday", "Sunday":
		println("It's the weekend.")
	default:
		println("Unknown day.")
	}

	// Using expressions
	number := 15
	switch {
	case number < 10:
		fmt.Println("Number is less than 10")
	case number >= 10 && number <= 20:
		fmt.Println("Number is between 10 and 20")
	default:
		fmt.Println("Number is greater than 20")
	}

	// Using fallthrough
	num := 2
	switch {
	case num > 1:
		fmt.Println("Greater than 1")
		fallthrough
	case num == 2:
		fmt.Println("Number is 2")
	default:
		fmt.Println("Not Two")
	}

	checkType(42)
	checkType(3.14)
	checkType("Hello")
	checkType(true) // Unknown Type

	// if else can be switched to switch case, if only 2 or 3 conditions are there, if more than 3 conditions are there, switch case is better.
}

func checkType(x interface{}) {
	// Type switch
	switch x.(type) {
	case int:
		fmt.Println("It's an integer")
		// fallthrough cannot be used in type switch
	case int32:
		fmt.Println("It's an int32")
	case float64:
		fmt.Println("It's a float")
	case string:
		fmt.Println("It's a string")
	default:
		fmt.Println("Unknown Type")
	}
}
