package main

import "fmt"

type Celsius float64

func (c Celsius) ToFahrenheit() float64 {
	return float64(c * 9 / 5 * 32)
}

type MyInt int

func main() {
	var num MyInt = 5
	fmt.Println(num)

	temp := Celsius(25)
	fmt.Println(temp.ToFahrenheit())
}

/* func main() {
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}
	fmt.Println("Finish loop")
} */

/* func main() {

	day := 1
	switch day {
	case 1:
		fmt.Println("Monday")
		fallthrough
	case 2:
		fmt.Println("Tuesday")
	default:
		fmt.Print("Other day")
	}
} */

/* func main() {
	numsMaps := map[string]int{"a": 1, "b": 2}

	for key, value := range numsMaps {
		fmt.Printf("Key: %s, Value: %d\n", key, value)
	}


	for i := 0; i <= 5; i++ {
		if i == 3 {
			break
		}
		fmt.Println(i)
	}

	for i := range 5 {
		if i%2 == 0 {
			continue
		}
		fmt.Println(i)
	}

	for i := range 3 {
		for j := range 3 {
			if j == 2 {
				continue
			}
			fmt.Printf("i=%d , j=%d\n", i, j)
		}

	}

	nums := []int{1, 2, 3, 4, 5}
	for i, num := range nums {
		fmt.Printf("Index: %d, Value: %d\n", i, num)
	}
} */

/* func main() {
	day := 1
	switch day {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	default:
		fmt.Print("Other day")
	}

	letter := "a"

	switch letter {
	case "a", "e", "i", "o", "u":
		fmt.Println("Vocal")
	default:
		fmt.Println("Consonant")
	}
} */

/* func sayingBye() string {
	return "Bye"
}

func main() {
	messageFunc := sayingBye

	println(messageFunc())

	fmt.Println(mathutil.Divide(6, 2))
} */

/* func main() {
	for i := 0; i < 5; i++ {
		fmt.Println(i)

	}

	nums := []int{1, 2, 3}
	for _, num := range nums {
		fmt.Println(num)
	}
} */

/* func main() {

	divide, err := mathutil.Divide(4, 2)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Result: ", divide)
	}

} */

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
