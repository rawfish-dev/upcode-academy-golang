package main

import (
	"fmt"
)

func main() {
	fileProcessor := FileProcessor{}

	err := fileProcessor.Process("some_non_existant_file.txt")
	PrintErrorType(err)

	err = fileProcessor.Process("emptyfile.txt")
	PrintErrorType(err)
}

func PrintErrorType(err error) {
	fmt.Printf("Error: %+v\n", err)

	/* Add the switch case for handling FileWasEmptyError */
	switch t := err.(type) {
	case FileCouldNotBeOpenedError:
		fmt.Println("Error is of type FileCouldNotBeOpenedError")
	default:
		fmt.Printf("\nError is of an unhandled type %+v \n", t)
	}
}
