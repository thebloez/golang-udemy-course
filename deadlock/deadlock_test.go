package deadlock

import (
	"fmt"
	"sync"
	"time"
)

type UserBalance struct {
	sync.Mutex
	Name    string
	Balance int
}

func (u *UserBalance) Lock() {
	u.Mutex.Lock()
}

func (u *UserBalance) Unlock() {
	u.Mutex.Unlock()
}

func (u *UserBalance) Change(amount int) {
	u.Balance += amount
}

func Transfer(user1 *UserBalance, user2 *UserBalance, amount int) {
	user1.Lock()
	fmt.Println("Lock", user1.Name)
	user1.Change(amount)

	time.Sleep(20 * time.Millisecond)

	user2.Lock()
	fmt.Println("Lock", user2.Name)
	user2.Change(amount)

	time.Sleep(20 * time.Millisecond)

	user1.Unlock()
	user2.Unlock()
}
