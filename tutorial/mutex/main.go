package account

import (
	"sync"
)

// Account represents a person's bank account
type Account struct {
	sync.RWMutex
	// RWMutex is a reader/writer lock, and can be held by multiple readers or a single writer. Both sync.Mutex and sync.RWMutex implements the sync.Locker interface, so will have both Lock and Unlock methods. But RWmutex has two additional methods: RLock and RUnlock.
	balance int64
	closed  bool
}

// Open creates a new account with some initial money
func Open(initialDeposit int64) *Account {
	if initialDeposit <= 0 {
		return nil
	}

	return &Account{balance: initialDeposit}
}

// Close shuts down an account by paying out the balance and return a boolean
func (acc *Account) Close() (payout int64, ok bool) {
	acc.Lock()
	defer acc.Unlock()
	// It is idiomatic in Go to call defer Unlock() immediately after Lock() method, so that programmers don’t forget this and cause a race condition. The defer keyword indicates that the execution of that function will occur when the surrounding function returns.
	if acc.closed {
		return 0, false
	}

	payout = acc.balance
	acc.balance = 0
	acc.closed = true
	return payout, true
}

// Balance outputs the current account balance
func (acc *Account) Balance() (balance int64, ok bool) {
	acc.Lock()
	defer acc.Unlock()
	if acc.closed {
		return 0, false
	}

	return acc.balance, true
}

// Deposit allows user to withdraw or deposit money into account
func (acc *Account) Deposit(balance int64) (newBalance int64, ok bool) {
	acc.Lock()
	defer acc.Unlock()
	if acc.closed {
		return 0, false
	}

	if acc.balance+balance < 0 {
		return acc.balance, false
	}

	acc.balance += balance
	return acc.balance, true
}
