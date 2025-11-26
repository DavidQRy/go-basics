package main

import (
	"fmt"
	"go-basics/mathutil"
)

func main() {

	divide, err := mathutil.Divide(4, 2)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Result: ", divide)
	}

}

/* var name = "David"

// const Pi, Height = 3.1416 , 1750
const (
	Pi     = 3.1416
	Height = 175
)

func main() {
	// fmt.Println("Hello word!")

	fmt.Println(mathutil.Add(5, 4))
	fmt.Println(mathutil.Sub(4, 3))

	fmt.Println(math.Sqrt(144))

	var x int = 5
	fmt.Println(x)

	var a, b = 3, 4

	fmt.Println(a, b)

	y, z := 5, 6

	fmt.Println(y, z)
	fmt.Println(name)
	fmt.Println(mathutil.Age)

	fmt.Println(Pi)

	fmt.Println(Height)
}
*/
