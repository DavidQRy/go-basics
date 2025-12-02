package main

import "fmt"

func main() {
	varMap := map[string]int{"apple": 2, "Banana": 3}
	fmt.Println(varMap["apple"])

	for key, value := range varMap {
		fmt.Printf("key: %s Value: %d n", key, value)
	}
}
