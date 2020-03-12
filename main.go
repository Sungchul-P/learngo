package main

import (
	"fmt"

	"github.com/Sungchul-P/learngo/mydict"
)

func main() {
	dictionary := mydict.Dictionary{}
	baseWord := "hello"
	dictionary.Add(baseWord, "First")
	err := dictionary.Update(baseWord, "Second")
	dictionary.Delete(baseWord)
	word, err := dictionary.Search(baseWord)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(word)
	}
}

// Account Example
/**
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
*/
