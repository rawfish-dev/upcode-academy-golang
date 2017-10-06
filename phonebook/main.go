package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
	display = "1"
	add     = "2"
	remove  = "3"
)

func main() {
	phoneBook := make(map[string]string)

	fmt.Println("Available Commands")
	fmt.Println("1) Display all entries in the phone book")
	fmt.Println("2) Add an entry to the phone book")
	fmt.Printf("3) Remove an entry from the phone book\n\n")

	for {
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Printf("Please choose either 1, 2 or 3: ")
		scanner.Scan()
		command := scanner.Text()

		switch command {
		case display:
			fmt.Printf("\n*********************\n")
			fmt.Println("* Kevin's Phone Book *")
			fmt.Printf("*********************\n\n")

			if len(phoneBook) == 0 {
				fmt.Printf("The phone book appears to be empty...\n\n")
				break
			}

			for name, number := range phoneBook {
				fmt.Printf("%s -> %s\n", name, number)
			}
			fmt.Println("")

		case add:
			fmt.Printf("Input name: ")
			scanner.Scan()
			name := scanner.Text()

			if name == "" {
				fmt.Println("You cannot add an empty name to the phone book :(")
				break
			}

			fmt.Printf("Input number: ")
			scanner.Scan()
			number := scanner.Text()

			if number == "" {
				fmt.Println("You cannot add an empty number to the phone book :(")
				break
			}

			phoneBook[name] = number

		case remove:
			if len(phoneBook) == 0 {
				fmt.Println("Oops! There is nothing to be removed currently.")
				break
			}

			fmt.Printf("Enter the name of the entry to be deleted: ")
			scanner.Scan()
			name := scanner.Text()

			if name == "" {
				fmt.Println("Oops! The name to be deleted cannot be empty.")
				break
			}

			_, ok := phoneBook[name]
			if !ok {
				fmt.Printf("Oops! %s was not found in the phone book.\n", name)
				break
			}

			delete(phoneBook, name)
		}
	}
}
