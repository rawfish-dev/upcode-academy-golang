package main

import (
	"fmt"
)

type X struct {
	Name string
	Y
	Z
}

type Y struct {
	Name string
}

type Z struct {
	Name string
}

func main() {
	x := X{
		Name: "123",
		Y:    Y{"abc"},
		Z:    Z{"def"},
	}
	fmt.Printf("X: %+v\n", x)
	fmt.Println(x.Name)
}
