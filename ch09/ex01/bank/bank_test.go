package bank_test

import (
	// "fmt"
	"testing"

	"github.com/relsa/exercise-gopl.io/ch09/ex01/bank"
)

// func TestBank(t *testing.T) {
// 	done := make(chan struct{})
//
// 	// Alice
// 	go func() {
// 		bank.Deposit(200)
// 		fmt.Println("=", bank.Balance())
// 		done <- struct{}{}
// 	}()
//
// 	// Bob
// 	go func() {
// 		bank.Deposit(100)
// 		done <- struct{}{}
// 	}()
//
// 	// Wait for both transactions.
// 	<-done
// 	<-done
//
// 	if got, want := bank.Balance(), 300; got != want {
// 		t.Errorf("Balance = %d, want %d", got, want)
// 	}
// }

func TestWithdraw(t *testing.T) {
	bank.Deposit(300)

	ts := []struct {
		amount  int
		rest    int
		success bool
	}{
		{200, 100, true},
		{100, 0, true},
		{50, 0, false},
	}

	for _, tc := range ts {
		ok := bank.Withdraw(tc.amount)
		if ok != tc.success {
			t.Errorf("got %v, want %v", ok, tc.success)
		}
		b := bank.Balance()
		if b != tc.rest {
			t.Errorf("balance %d, want %d", b, tc.rest)
		}
	}
}
