package main

import (
	"fmt"
	"sync"
	"time"
)

type Bank struct {
	balance int
}

func deposit(bank *Bank) {
	for i := 0; i < 10; i++ {
		bal := bank.balance
		fmt.Println("in deposit - Current Balance is ", bal)
		time.Sleep(10 * time.Millisecond)
		bal = bal + 1
		bank.balance = bal
	}
}
func widraw(bank *Bank) {
	for i := 0; i < 10; i++ {
		bal := bank.balance
		fmt.Println("in widraw - Current Balance is ", bal)
		time.Sleep(10 * time.Millisecond)
		bal = bal - 1
		bank.balance = bal
	}
}
func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	bank := Bank{0}

	//anonymous function call in go routine.
	go func() {
		deposit(&bank)
		wg.Done()
	}()
	go func() {
		widraw(&bank)
		wg.Done()
	}()

	fmt.Println(" a line before wait ...", wg)

	wg.Wait()
	fmt.Println(" a line after wait ...")
	fmt.Println("Current balance = ", bank.balance)

}
