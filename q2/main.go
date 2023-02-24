package main

import (
	"fmt"
	"time"
)

func importMessage(input chan string) {
	input <- "Belize is going to blow up in 45 minutes."
}

func displayMessage(input chan string) {
	fmt.Println("Important announcement:", <-input)
}

func main() {
	message := make(chan string)

	go importMessage(message)
	go displayMessage(message)

	time.Sleep(1 * time.Second)
}
