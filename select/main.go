package main

import (
	"fmt"
	"time"
)

func main() {

	ch := make(chan string)

	go func() {
		time.Sleep((2 * time.Second))
		ch <- "Done"
	}()

	select {
	case msg := <-ch:
		fmt.Println(msg)
	case <-time.After(1 * time.Second):
		fmt.Println("Timeout: No message received within 1 second")
	}
}

/* func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		ch1 <- "Channel 1"
	}()

	go func() {
		ch2 <- "Channel 2"
	}()

	select {
	case msg1 := <-ch1:
		fmt.Println(msg1)
	case msg2 := <-ch2:
		fmt.Println(msg2)
	}
} */
