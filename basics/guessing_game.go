package main

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
	fmt.Println("Welcome to the guessing game!")
	fmt.Println("I have chosen a number between 1 and 100.")
	fmt.Println("Can you guess what it is?")

	var guess int
	for {
		fmt.Print("Enter your guess: ")
		fmt.Scanln(&guess)

		if guess == target {
			fmt.Println("Congratulations! You guessed the correct number!")
			break
		} else if guess < target {
			fmt.Println("Too low! Try guessing a higher number.")
		} else {
			fmt.Println("Too high! Try guessing a lower number.")
		}
	}
}
