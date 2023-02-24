package main

import (
	"fmt"
	"sync"
)

func importMessage(input chan string) {
	input <- "Belize is going to blow up in 45 minutes."
}

func displayMessage(input chan string) {
	fmt.Println("Important announcement:", <-input)
}

func main() {
	message := make(chan string)
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		importMessage(message)
		wg.Done()
	}()

	go func() {
		displayMessage(message)
		wg.Done()
	}()
	wg.Wait()
}
