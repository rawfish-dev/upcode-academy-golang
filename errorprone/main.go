package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 10; i++ {
		/* Make some calls to EvenErrors and handle errors accordingly */
	}

	Execute()
}

func EvenErrors(value int) error {
	if value%2 == 0 {
		/* The fmt package offers a way to create a formatted error
		look up the documentation and try to find it! */
		return
	}

	return nil
}

func Execute() {
	/* Add a defer and recover function to catch the panic */

	/* Move the existing panic and defer functions around to get the correct output */
	panic("BOOM")

	defer func() {
		defer func() {
			fmt.Println("some deferred function")
		}()

		fmt.Println("another deferred function")
	}()
}
