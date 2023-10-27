package main

import (
	"fmt"
	"time"
)

func printMessage(text string, channel chan string)  {
	for i :=0; i < 5; i++ {
		fmt.Println(text)
		time.Sleep(800 * time.Millisecond)
	}

	// channels can only be assigned by using the <- operator
	// done is sent to the channel variable
	channel <- "Done!!"
}

// GOROUTINES
// kind of like threads in other languages


// CHANNELS
// this enables different goroutines to communicate with each other
// this contains a value of any kind which can be accessed and waited for in other goroutines

// so they can be used to send data from one goroutine to the other
// and they can be used to make a goroutine wait 

// this is executing in the main goroutine
// the app ends when func main ends
func main() {
	// the chan keyword creates a channel variable
	var channel chan string
	// you can open more go routines by calling go
	go printMessage("Go is great", channel)
	// printMessage("We love Go")


	// we can delay a func by making them wait for the channel variable
	response := <- channel
	close(channel)
	fmt.Println(response)


	// logs := make(chan string, 2)

	// logs <- "hello"
	// logs <- "world"

	// fmt.Println(logs)
}
