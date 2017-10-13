package main

import (
	"fmt"
)

func main() {
	Execute()
}

func Execute() {
	defer DoThis()

	/* Defer the DoThat function here and observe the order! */

	/* Add an anonymous function as part of the defer chain */

	fmt.Println("executing!")
}

func DoThis() {
	fmt.Println("doing this")
}

func DoThat() {
	fmt.Println("or doing that")
}
