package bank

type wreq struct {
	amount  int
	succeed chan bool
}

var deposits = make(chan int) // send amount to deposit
var balances = make(chan int) // receive balance
var withdraws = make(chan *wreq)

func Deposit(amount int) { deposits <- amount }
func Balance() int       { return <-balances }
func Withdraw(amount int) bool {
	succ := make(chan bool)
	withdraws <- &wreq{amount, succ}
	return <-succ
}

func teller() {
	var balance int // balance is confined to teller goroutine
	for {
		select {
		case amount := <-deposits:
			balance += amount
		case balances <- balance:
		case withdraw := <-withdraws:
			if balance < withdraw.amount {
				withdraw.succeed <- false
			} else {
				balance -= withdraw.amount
				withdraw.succeed <- true
			}
		}
	}
}

func init() {
	go teller() // start the monitor goroutine
}
