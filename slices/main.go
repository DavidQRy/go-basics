package main

import (
	"fmt"
)

func main() {
	arr := [5]int{10, 20, 30, 40, 50}
	slice := arr[1:4]
	fmt.Println(slice)
	fmt.Println(arr)

}

/* func main() {
	slice := []int{1, 2, 3}
	slice = append(slice, 4, 5)
	fmt.Println(slice)
} */
