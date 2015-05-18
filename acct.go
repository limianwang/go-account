package acct

import "strconv"

// Account account object
type Account struct {
	ID  int
	Bal float64
}

// Balance returns balance of the account
func (a *Account) Balance() float64 {
	val, _ := strconv.ParseFloat(strconv.FormatFloat(a.Bal, 'f', 2, 64), 64)
	return val
}

// Transaction transaction object
type Transaction struct {
	operations []Operation
}

// Operation operation that will occur on a MoveMoney
type Operation struct {
	amount  float64
	fromAcc *Account
	toAcc   *Account
}

// MoveMoney prepares the money flowing from one to other
func (t *Transaction) MoveMoney(p float64, fromAcc *Account, toAcc *Account) {
	t.prepare(p, fromAcc, toAcc)
}

// Close commits the operations
func (t *Transaction) Close() {
	if err := t.commit(); err != nil {
		t.rollback()
	}
}

func (t *Transaction) prepare(p float64, fromAcc *Account, toAcc *Account) {
	op := Operation{}
	op.amount = p
	op.fromAcc = fromAcc
	op.toAcc = toAcc
	t.operations = append(t.operations, op)
}

func (t *Transaction) commit() error {
	for _, opt := range t.operations {
		if opt.amount == 0 {
			continue
		} else {
			opt.fromAcc.Bal -= opt.amount
			opt.toAcc.Bal += opt.amount
		}
	}
	return nil
}

func (t *Transaction) rollback() error {
	return nil
}

// NewTransaction returns a new transaction
func NewTransaction() *Transaction {
	t := &Transaction{}
	t.operations = make([]Operation, 5)
	return t
}

// NewAccount returns a new account
func NewAccount() *Account {
	return &Account{}
}
