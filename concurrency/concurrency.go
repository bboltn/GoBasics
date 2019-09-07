package main

import (
	"fmt"
	"time"
)

type character struct {
	name string
}

func sendCharacters(location chan character, cs ...character) {
	for _, c := range cs {
		println("Sending", c.name)
		location <- c
	}

	// once we're done sending characters,
	// let's close the channel to tell the receiver to stop looking for values
	close(location)
}

func receiveCharacters(location chan character) {
	// as characters come in on the channel, let's output their value
	for c := range location {
		fmt.Println("Received!", c.name)
	}
}

func main() {
	akira := character{"Akira"}
	tetsuo := character{"Tetsuo"}
	kei := character{"kei"}

	olympicStadium := make(chan character)
	miyakoTemple := make(chan character)

	// goroutine is a light weight thread
	go sendCharacters(olympicStadium, akira, tetsuo)
	go sendCharacters(miyakoTemple, kei)

	receiveCharacters(miyakoTemple)
	receiveCharacters(olympicStadium)

	// buffered channels allow you to send a certain number
	// of items before it deadlocks
	bufferedChannel := make(chan character, 2)
	bufferedChannel <- kei
	bufferedChannel <- tetsuo
	fmt.Println(<-bufferedChannel) // comment this out and you'll get a deadlock error
	bufferedChannel <- tetsuo

	// use a select statement to wait on multiple channels
	c1, c2, c3, done := make(chan character), make(chan character), make(chan character), make(chan character)

	//start up our goroutine to add characters to the channels
	go func() {
		println("Sending characters...")
		c2 <- character{"akira"}

		println("Pausing")
		time.Sleep(5 * time.Second)
		c3 <- character{"tetsuo"}
		c1 <- character{"kei"}
		done <- character{"done"}
	}()

	// here we will block until we get characters on the channel
	// in the goroutine above, we pause for 5 seconds.  You'll see
	// that we get Akira, wait, and then get the next characters.
	for {
		select {
		case v := <-c1:
			fmt.Println("Channel 1:", v.name)
		case v := <-c2:
			fmt.Println("Channel 2:", v.name)
		case v := <-c3:
			fmt.Println("Channel 3:", v.name)
		case <-done:
			fmt.Println("quit!")
			return
		}
	}
}
