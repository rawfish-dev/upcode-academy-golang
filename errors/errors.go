package main

import (
	"errors"
	"fmt"
)

func main() {
	err := createError(true)
	printError(err)

	err = createError(false)
	/* Add some code here to prevent the program from crashing */
	printError(err)

	/* Enhance the program by uncommenting out these */
	/* lines and filling in the missing parts */
	// err = createCustomError()
	// printError(err)
}

func createError(control bool) error {
	if control {
		return errors.New("An error was asked to be returned!")
	}

	return nil
}

func printError(err error) {
	fmt.Println(err.Error())
}

/* Create a new struct and make it satisfy the Error interface */

/* Create the createCustomError function that creates your custom error type */
