package main

import (
	"fmt"
)

func main() {
	/*
			Try moving the defer and recover block here.
		   	Notice anything different in the output?
	*/

	fmt.Println("Before the panic")

	panicNow("Causing a panic!")

	fmt.Println("After the panic")
}

func panicNow(message string) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in panicNow", r)
		}
	}()

	panic("crashes the program")
}
