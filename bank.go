package main

import(
	"fmt"
	"strings"
	"math/rand"
)

type Account struct {
	number string
	owner Entity
	balance float64
	interestRate float64
	accountType string
}

type Entity struct {
	id string
	address string
	entityType string
}

type Wallet struct {
	walletId string
	accounts[] Account
	owner Entity
}

func withdraw(a Account, amount float64) bool {
	if a.balance < amount {
		fmt.Println("You do not have enough money in your bank account.")
		return false
	} else {
		a.balance -= amount
		return true
	}
}

func deposit(a Account, amount float64) {
	a.balance += amount
}

func applyInterest(a Account) {
	if a.owner.entityType == "Individual" {
		if a.accountType == "Checking" {
			a.interestRate = 1
		} else if a.accountType == "Investment" {
			a.interestRate = 2
		} else if a.accountType == "Saving" {
			a.interestRate = 5
		}
	} else if a.owner.entityType == "Business" {
		if a.accountType == "Checking" {
			a.interestRate = 0.5
		} else if a.accountType == "Investment" {
			a.interestRate = 1
		} else if a.accountType == "Saving" {
			a.interestRate = 2
		}
	}
}

func wire(source Account, target Account, amount float64) {
	if withdraw(source, amount) {
		deposit(target, amount)
	}
}

func changeAddress(e Entity, newAddress string) {
	e.address = newAddress
}

func displayAccounts(w Wallet) {
	for i := 0; i < len(w.accounts); i++ {
		if w.accounts[i].accountType == "Checking" {
			fmt.Println(w.accounts[i])
		}
	}
	for i := 0; i < len(w.accounts); i++ {
		if w.accounts[i].accountType == "Investment" {
			fmt.Println(w.accounts[i])
		}
	}
	for i := 0; i < len(w.accounts); i++ {
		if w.accounts[i].accountType == "Saving" {
			fmt.Println(w.accounts[i])
		}
	}
}

func balance(w Wallet) {
	for i := 0; i < len(w.accounts); i++ {
		fmt.Println(w.accounts[i].balance)
	}
}

func wire(w Wallet, source Account, target Account, amount float64) {
	included := false
	for i := 0; i < len(w.accounts); i++ {
		if w.accounts[i] == source {
			included = true
			wire(source, target, amount)
			break
		}
	}
	if (!included) {
		fmt.Println("Source account is not included in the wallet.")
	}
}

func main() {

}