package basics

import "fmt"

func main() {
	message := "Hello World"
	for i, v := range message {
		fmt.Println(i, v)                         // print unicode code point, type of v is 'rune', from int32
		fmt.Printf("Index: %d, Rune: %c\n", i, v) // print unicode value
	}
}
