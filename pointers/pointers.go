package pointers

import (
	"errors"
	"fmt"
)

// Error that is raised when more bitcoins are withdrawn than in the wallet
var ErrNotEnoughBitcoins = errors.New("Not enough Bitcoins in your wallet")

// NewType for Bitcoin, based on int
type Bitcoin int

// Stringer for Bitcoins
func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

type Wallet struct {
	balance Bitcoin
}

// Add Value to wallet
func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

// Remove value from wallet
func (w *Wallet) Withdraw(amount Bitcoin) error {
	if amount > w.balance {
		return ErrNotEnoughBitcoins
	}

	w.balance -= amount
	return nil
}

// Get the current balance from the wallet
func (w *Wallet) Balance() Bitcoin {
	return w.balance
}
