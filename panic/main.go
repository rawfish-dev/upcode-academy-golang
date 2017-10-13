package main

import (
	"fmt"
)

func main() {
	fmt.Println("Before the panic")

	panicNow("Causing a panic!")

	fmt.Println("After the panic")
}

func panicNow(message string) {
	panic("crashes the program")
}
