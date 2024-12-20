package bitcoin

import (
	"errors"
	"fmt"
)

var errorInsufficientFunds = errors.New("cannot withdraw, insufficient funds")

type Stringer interface {
	String() string
}
type Bitcoin int
type Wallet struct {
	balance Bitcoin
}

func (w *Wallet) Deposite(amount Bitcoin) {

	w.balance += amount
}
func (w *Wallet) Balance() Bitcoin {
	return w.balance
}
func (w *Wallet) Withdraw(amount Bitcoin) error {
	if amount > w.balance {
		return errorInsufficientFunds
	}

	w.balance -= amount

	return nil
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}
