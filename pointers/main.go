package main

import "fmt"

type bankAccount struct {
	Landlord string
	Balance  float64
}

func (ba bankAccount) DepositByValue(amount float64) {
	ba.Balance += amount
	fmt.Printf("Balance after deposit (method by value): %.2f\n", ba.Balance)
}

func (ba *bankAccount) Deposit(quantity float64) {
	ba.Balance += quantity
	fmt.Printf("Balance after of deposit: %.2f \n", ba.Balance)
}

func (ba *bankAccount) Withdraw(quantity float64) {
	if ba.Balance >= quantity {
		ba.Balance -= quantity
		fmt.Printf("successful withdrawal, current balance: %.2f\n", ba.Balance)
	} else {
		fmt.Println("insuffient funds")
	}
}

func main() {
	account := bankAccount{
		Landlord: "David Quintero",
		Balance:  2000.0,
	}

	fmt.Printf("beginning balance: %.2f\n", account.Balance)

	account.DepositByValue(100)
	fmt.Printf("balance after of depositByValue: %v\n", account.Balance)

	account.Deposit(500)
	fmt.Printf("Rea balance after of deposit %.2f \n", account.Balance)

	account.Withdraw(450)
	fmt.Printf("final Balance: %.2f\n", account.Balance)
	fmt.Println(account.Landlord)
}

/* type LargeStruct struct {
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
} */

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
