package main

import "fmt"

func printAnything(v interface{}) {
	fmt.Println(v)
}

func main() {
	printAnything(42)
	printAnything("Hi!")
	printAnything(3.14)
}

/* type Talkative interface {
	Talk() string
}

type Dog struct{}

func (d Dog) Talk() string {
	return "Wuuff!!"
}

func main() {
	var dog Talkative = Dog{}

	fmt.Println(dog.Talk())
} */
