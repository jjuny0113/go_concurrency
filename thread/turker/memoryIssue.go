package turker

import (
	"fmt"
	"sync"
	"time"
)

type Account struct {
	Balance int
}

func depositAndWithdraw(account *Account) {
	if account.Balance < 0 {
		panic(fmt.Sprintf("Balance Should not be negative value: %d", account.Balance))

	}
	account.Balance += 1000
	time.Sleep(time.Millisecond)
	account.Balance -= 1000

}

func MemoryIssue() {
	var wg sync.WaitGroup

	account := &Account{0}    // ❶ 0원 잔고 통장
	wg.Add(10)                // ❷ WaitGroup 객체 생성
	for i := 0; i < 10; i++ { // ❸ Go 루틴 10개 생성
		go func() {
			for {
				depositAndWithdraw(account)
			}
			wg.Done()
		}()
	}
	wg.Wait()
}

var mutex sync.Mutex

func depositAndWithdrawV2(account *Account) {
	mutex.Lock()
	defer mutex.Unlock()
	if account.Balance < 0 {
		panic(fmt.Sprintf("Balance Should not be negative value: %d", account.Balance))

	}
	account.Balance += 1000
	time.Sleep(time.Millisecond)
	account.Balance -= 1000

}

func MemoryIssueSolution() {
	var wg sync.WaitGroup

	account := &Account{0}    // ❶ 0원 잔고 통장
	wg.Add(10)                // ❷ WaitGroup 객체 생성
	for i := 0; i < 10; i++ { // ❸ Go 루틴 10개 생성
		go func() {
			for {
				depositAndWithdraw(account)
			}
			wg.Done()
		}()
	}
	wg.Wait()
}
