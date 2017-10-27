package main

import (
	"fmt"
	"strconv"
	"time"
)

func read(inputChan, outputChan chan string) {
	for {
		select {
		case input, ok := <-inputChan:
			if !ok {
				fmt.Println("read close")
				return
			}

			fmt.Println("read input", input)
			outputChan <- input
		}
	}
}

func process(inputChan, outputChan chan string) {
	for {
		select {
		case input, ok := <-inputChan:
			if !ok {
				fmt.Println("process close")
				return
			}

			fmt.Println("process input", input)
			outputChan <- input
		}
	}
}

func output(inputChan, outputChan chan string) {
	for {
		select {
		case input, ok := <-inputChan:
			if !ok {
				fmt.Println("output close")
				return
			}

			fmt.Println("output input", input)
			outputChan <- input
		}
	}
}

func main() {
	readInputChan := make(chan string)
	processInputChan := make(chan string)
	outputInputChan := make(chan string)
	finalOutputChan := make(chan string)

	go read(readInputChan, processInputChan)
	go process(processInputChan, outputInputChan)
	go output(outputInputChan, finalOutputChan)

	const iterations = 3

	for i := 0; i < iterations; i++ {
		go func(iterationCount int) {
			readInputChan <- strconv.Itoa(iterationCount)
		}(i)
	}

	outputCount := 0
	for {
		select {
		case finalOutput := <-finalOutputChan:
			fmt.Println("received", finalOutput)
			outputCount++
		}

		if outputCount >= iterations {
			break
		}
	}

	close(readInputChan)
	close(processInputChan)
	close(outputInputChan)

	// Wait for cleanup to occur
	<-time.After(time.Second)

	fmt.Println("done!")
}
