package main

import (
	"fmt"
)

func main() {
	dataChannel := make(chan string) // create a readable and writable channel

	dataChannel <- "a string value" // write to channel

	value := <-dataChannel // read from channel
	fmt.Println(value)

	close(dataChannel) // close the channel
}
