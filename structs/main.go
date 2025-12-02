package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func main() {
	person := Person{Name: "David", Age: 19}
	name := person.Name
	fmt.Println(name)
	fmt.Println(person)
}
