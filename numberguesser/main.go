package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

const (
	maxTries    = 5
	minBoundary = 0
	maxBoundary = 100
)

func main() {
	lower := minBoundary
	upper := maxBoundary

	answer := generateRandomNumber(maxBoundary)

	fmt.Printf("Game is starting! Please choose an integer between %d to %d\n", minBoundary, maxBoundary)

	for i := 0; i < maxTries; i++ {
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Printf("Please enter your guess: ")
		scanner.Scan()
		enteredGuess := scanner.Text()

		guessValue, err := strconv.Atoi(enteredGuess)
		if err != nil {
			panic(err.Error())
		}

		// Validate entered value
		validInput := validate(guessValue)
		if !validInput {
			fmt.Printf("Please enter a number between %d and %d only! %d attempt(s) left...\n", minBoundary, maxBoundary, maxTries-1-i)
			continue
		}

		// Check the user has guessed correctly
		if guessValue == answer {
			fmt.Println("Correct!")
			os.Exit(0)
		}

		// Shift the boundaries accordingly
		if guessValue < answer {
			lower = guessValue
		} else {
			upper = guessValue
		}

		fmt.Printf("%d attempt(s) left... Please choose an integer between %d to %d\n", maxTries-1-i, lower, upper)
	}

	// No guesses were correct
	fmt.Printf("Game over, the answer was %d. Better luck next time!\n", answer)
}

func generateRandomNumber(ceiling int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(ceiling + 1)
}

func validate(input int) bool {
	return input >= minBoundary && input <= maxBoundary
}
