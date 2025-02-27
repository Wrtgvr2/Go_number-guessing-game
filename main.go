package main

import (
	"fmt"
	"math/rand"
)

var ranNumber int
var attempts int

func main() {
	currentAttempt := 0
	playerWon := false

	fmt.Println("Welcome to the Number Guessing Game!")
	gameInit()
	fmt.Println("I'm thinking of a number between 1 and 100.")
	fmt.Printf("You have %d chances to guess the correct number.\n", attempts)

gameLoop:
	for currentAttempt < attempts {
		var guess int

		currentAttempt += 1

		fmt.Print("Enter your guess: ")
		fmt.Scan(&guess)

		switch {
		case ranNumber < guess:
			fmt.Printf("Incorrect! The number is less than %d\n", guess)
		case ranNumber > guess:
			fmt.Printf("Incorrect! The number is greater than %d\n", guess)
		case ranNumber == guess:
			playerWon = true
			break gameLoop
		}
	}

	if playerWon {
		fmt.Printf("\nCongratulations! You guessed the correct number in %d attempts.", currentAttempt)
	} else {
		fmt.Printf("\nYou've run out of attempts, better luck next time! Correct number is %d", ranNumber)
	}
	fmt.Println("\nPress Enter to exit...")
	fmt.Scanln()
	fmt.Scanln()
}

func gameInit() {
	chooseDifficulty()
	generateNumberToGuess()
}

func chooseDifficulty() {
	var diff int
	var stopLoop bool = false

	for !stopLoop {
		fmt.Print("Please select the difficulty level:\n1. Easy (10 chances)\n2. Medium (5 chances\n3. Hard (3 chances)\nDifficulty: ")
		fmt.Scanln(&diff)

		switch diff {
		case 1:
			attempts = 10
			stopLoop = true
		case 2:
			attempts = 5
			stopLoop = true
		case 3:
			attempts = 3
			stopLoop = true
		}
	}
}

func generateNumberToGuess() {
	ranNumber = rand.Intn(99) + 1
}
