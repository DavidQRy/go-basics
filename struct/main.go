package main

import "fmt"

type Rectangle struct {
	Width, Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func main() {
	rect := Rectangle{
		Width:  4,
		Height: 5,
	}
	fmt.Println(rect.Area())
}

/* type People struct {
	Name string
	Age  int
}

func main() {
	people := People{
		Name: "David",
		Age:  19,
	}

	fmt.Println(people)
} */
