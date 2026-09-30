package basics

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	source := rand.NewSource(time.Now().UnixNano())
	random := rand.New(source)

	// Generate a random number between 1 and 100
	target := random.Intn(100) + 1

	// Welcome message
	fmt.Println("Welcome to the Guessing Game!")
	fmt.Println("I have chosen a number between 1 and 100")
	fmt.Println("Can you guess what it is?")

	var guess int
	for {
		fmt.Println("Enter your guess: ")
		// The copy of guess is passed to Scanln, so the value of guess will be updated with the user input, and the original variable guess is not updated.
		// fmt.Scanln(guess)

		// The address of guess is passed to Scanln, so the value of guess will be updated with the user input, and the original variable guess is updated. '&' is used to get the memory address of the variable guess, so that Scanln can modify the value of guess directly.
		fmt.Scanln(&guess)

		// Check if the guess is correct
		if guess == target {
			fmt.Println("Congratulations! You guessed the correct number:", target)
			break
		} else if guess < target {
			fmt.Println("Too low! Try again.")
		} else {
			fmt.Println("Too high! Try again.")
		}
	}

}
