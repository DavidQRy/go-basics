package main

import "fmt"

type LargeStruct struct {
	Data [1000]int
	Name string
	ID   int
}

func processByValue(ls LargeStruct) {
	fmt.Println("Processing: ", ls.Name)
}

func processByPointer(ls *LargeStruct) {
	fmt.Println("Processing: ", ls.Name)
}

func main() {
	large := LargeStruct{
		Name: "David",
		ID:   1,
	}

	processByValue(large)
	processByPointer(&large)
}

/* func increaseByValue(x int) {
	x = x + 1
	fmt.Println("Within the function", x)
}

func increaseByPointer(x *int) {
	*x = *x + 1
	fmt.Println("Within the function", *x)

}

func main() {
	num := 10

	increaseByValue(num)
	fmt.Println("after of increaseByValue: ", num)

	increaseByPointer(&num)
	fmt.Println("after of increaseByPointer: ", num)
} */

/* func main() {
	x := 20
	p := &x
	*p = 50

	fmt.Println(x)
} */

/* func main() {
	x := 10
	p := &x

	fmt.Println(p)
	fmt.Println(*p)
}
*/
