package main

import (
	"fmt"

	"github.com/Sungchul-P/learngo/accounts"
)

func main() {
	account := accounts.NewAccount("devnori")
	account.Deposit(10)
	account.ChangeOwner("ch-devnori")
	// fmt.Println(account.Balance())
	// err := account.Withdraw(20)
	// if err != nil {
	// 	// log.Fatalln(err)
	// 	fmt.Println(err)
	// }
	// fmt.Println(account.Balance(), account.Owner())
	fmt.Println(account)
}
