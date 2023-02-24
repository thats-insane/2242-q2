package main

import (
	"fmt"
	"sync"
)

func main() {
	message := make(chan string)
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		message <- "Belize is going to blow up in 45 minutes."
		wg.Done()
	}()

	go func() {
		fmt.Println("Important announcement:", <-message)
		wg.Done()
	}()
}
