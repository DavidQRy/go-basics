package main

import "fmt"

func main() {
	i := 0
Loop:
	if i < 3 {
		fmt.Println(i)
		i++
		goto Loop
	}
}

/* func main() {
	fmt.Println("Before of goto")
	goto End
	fmt.Println("This will not be executed")
End:
	fmt.Println("After goto")
} */
