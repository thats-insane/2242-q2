package main

import (
	"fmt"
	"time"
)

func displayMessage(input string) {
	fmt.Println("Important announcement:", input)
}

func main() {
	message := "Belize is going to blow up in 45 minutes."

	go displayMessage(message)

	time.Sleep(1 * time.Second)
}
