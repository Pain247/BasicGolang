package main

import (
	"fmt"
)

type Account struct {
	name   string
	amount float32
}
type Hi struct {
	name string
}

//Chỉ được sử dụng bởi type Account (coi như quy định method cho struct)
func (a *Account) setName(name string) {
	a.name = name
}

func (a *Account) getAmount() float32 {
	return a.amount
}
func (a *Account) addFund(number float32) {
	a.amount += number
}
func (a *Account) withdrawMoney(number float32) bool {
	if a.amount-number > 50.0 {
		a.amount -= number
		return true
	} else {
		return false
	}

}
func (a *Account) print() {
	fmt.Printf("Name : %s\n", a.name)
	fmt.Printf("Amount: %g", a.amount)
}
func (a *Hi) print() {
	fmt.Printf("Name : %s\n", a.name)
}
func main() {
	var t Hi
	t.print()
}
