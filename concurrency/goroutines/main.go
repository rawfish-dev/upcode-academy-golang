package main

import (
	"fmt"
	"time"
)

func printAfterDelay(iteration int) {
	<-time.After(1 * time.Second)
	fmt.Println(iteration)
}

func main() {
	for i := 0; i < 10; i++ {
		go printAfterDelay(i)
	}

	<-time.After(2 * time.Second)

	fmt.Println("OK!")
}
