package main

import "fmt"

func main() {
	arr := [3]string{"Go", "JavaScript", "php"}
	fmt.Println("My languages: ")

	for i, v := range arr {
		fmt.Printf("Index: %d Value: %s", i, v)
		fmt.Println()
	}
}

/* func main() {
	var arr [5]int
	arr[0] = 10
	arr[1] = 20

	fmt.Println(arr)
	fmt.Println(arr[1])
} */
