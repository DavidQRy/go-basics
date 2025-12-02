package main

import "fmt"

func main() {
	mp := make(map[string]int)
	mp["apple"] = 2
	mp["banana"] = 4

	fmt.Println(mp["apple"])
}
