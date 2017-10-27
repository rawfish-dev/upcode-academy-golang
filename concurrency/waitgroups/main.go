package main

import (
	"fmt"
	"sync"
	"time"
)

func printAfterDelay(iteration int) {
	<-time.After(1 * time.Second)
	fmt.Println(iteration)
}

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(iteration int) {
			printAfterDelay(iteration)
			wg.Done()
		}(i)
	}

	wg.Wait()

	fmt.Println("OK!")

}
