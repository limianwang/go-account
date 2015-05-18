package acct

import (
	"strconv"
)

type Account struct {
	ID  int
	Bal float64
}

func (a *Account) Balance() float64 {
	val, _ := strconv.ParseFloat(strconv.FormatFloat(a.Bal, 'f', 2, 64), 64)
	return val
}

type Transaction struct {
	operations []Operation
}

type Operation struct {
	amount  float64
	fromAcc *Account
	toAcc   *Account
}

func (t *Transaction) MoveMoney(p float64, fromAcc *Account, toAcc *Account) {
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

func (t *Transaction) Close() {
	if err := t.commit(); err != nil {
		t.rollback()
	}
}

func NewTransaction() *Transaction {
	t := &Transaction{}
	t.operations = make([]Operation, 5)
	return t
}

func NewAccount() *Account {
	return &Account{}
}
