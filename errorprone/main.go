package main

import (
	"fmt"
)

/* Output:
0 is even, so here's an error
2 is even, so here's an error
4 is even, so here's an error
6 is even, so here's an error
8 is even, so here's an error
Recovered in Execute: BOOM
some defered function
another defered function
*/

func main() {
	defer func() {
		fmt.Println("another deferred function")
	}()

	defer func() {
		fmt.Println("some deferred function")
	}()

	for i := 0; i < 10; i++ {
		/* Make some calls to EvenErrors and handle errors accordingly */
		err := EvenErrors(i)
		if err != nil {
			fmt.Println(err.Error())
		}
	}

	Execute()
}

func EvenErrors(value int) (err error) {
	if value%2 == 0 {
		/* The fmt package offers a way to create a formatted error
		look up the documentation and try to find it! */
		err := fmt.Errorf("%d is even, so here's an error", value)
		return err
	}

	return nil
}

func Execute() {
	/* Add a defer and recover function to catch the panic */
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in Execute:", r)
		}
	}()

	/* Move the existing panic and defer functions around to get the correct output */
	panic("BOOM")
}
