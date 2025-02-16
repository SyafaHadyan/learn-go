package main

import (
	"errors"
	"fmt"
)

type Money int

type Wallet struct{
    balance Money
}

type Stringer interface {
    String() string
}

func (wallet *Wallet) Deposit(amount Money) {
    wallet.balance += amount
}

func (wallet *Wallet) GetBalance() Money {
    return wallet.balance
}

func (wallet *Wallet) Withdraw(amount Money) error {
    if amount < wallet.balance{
        wallet.balance -= amount;
        return nil
    } else {
        return errors.New("Withdraw amount exceeds balance")
    }
}

func (money Money) String() string {
    return fmt.Sprintf("%d Money", money)
}

func main() {
    var wallet = Wallet{}
    wallet.Deposit(10)
    fmt.Println(wallet.GetBalance())
    // fmt.Println(wallet.String())
}
