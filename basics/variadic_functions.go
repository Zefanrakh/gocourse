package basics

import "fmt"

func main() {
	// ... Ellipsis
	// func functionName(param1 type1, param2 type2, param3 ...type3) returnType {

	// }

	fmt.Println(sum("sum of 1, 2, 3:", 1, 2, 3))

	// Using slices
	numbers := []int{1, 2, 3, 4, 5}
	fmt.Println(sum("sum of these numbers:", numbers...))

}

// Variadic parameters, must be the last
func sum(returnString string, nums ...int) (statement string, total int) {
	total = 0
	for _, v := range nums {
		total += v
	}
	statement = returnString
	return
}
