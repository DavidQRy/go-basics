package main

import (
	"fmt"
	"time"
)

func printNumbers(num int) {
	for i := 0; i <= num; i++ {
		fmt.Println(i)
		time.Sleep(100 * time.Millisecond)
	}
}

func main() {
	go printNumbers(5)
	go printNumbers(5)

	time.Sleep(time.Second)
}

/* func greet() {
	fmt.Println("Hello!")
}
func main() {
	go greet()

	time.Sleep(time.Second)

	fmt.Println("Main function completed")
}
*/
