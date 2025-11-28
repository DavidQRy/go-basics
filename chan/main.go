package main

import "fmt"

func main() {
	ch := make(chan string, 2)
	ch <- "Hello"
	ch <- "World"

	fmt.Println(<-ch)
	fmt.Println(<-ch)
}

/* func main() {
	ch := make(chan int)
	go func() {
		ch <- 42
	}()
	value := <-ch
	fmt.Println(value)
} */
