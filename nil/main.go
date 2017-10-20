package main

import (
	"fmt"
)

type SomeInterface interface {
}

type Something struct {
}

func main() {
	var someMap map[string]string
	if someMap == nil {
		fmt.Println("map is nil!")
	}

	var someSlice []int
	if someSlice == nil {
		fmt.Println("slice is nil!")
	}

	var someStruct *Something
	if someStruct == nil {
		fmt.Println("someStruct is nil!")
	}

	var someInterface SomeInterface
	if someInterface == nil {
		fmt.Println("someInterface is nil!")
	}
}
