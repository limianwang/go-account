package acct

import (
	"database/sql"

	// Import
	_ "github.com/go-sql-driver/mysql"
)

// Account account object
type Account struct {
	db  *sql.DB
	ID  int64
	Bal float64
}

// Balance returns balance of the account
func (a *Account) Balance() float64 {
	err := a.db.QueryRow("SELECT balance FROM account WHERE id = ?", a.ID).Scan(&a.Bal)
	if err != nil {
		panic(err)
	}
	return a.Bal

}

// Transaction transaction object
type Transaction struct {
	db *sql.DB
	tx *sql.Tx
}

// Begin initiates txn
func (t *Transaction) Begin() {
	t.tx, _ = t.db.Begin()
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
	t.tx.Exec("UPDATE account SET balance = balance - ? WHERE id = ?", p, fromAcc.ID)
	t.tx.Exec("UPDATE account SET balance = balance + ? WHERE id = ?", p, toAcc.ID)
}

func (t *Transaction) commit() error {
	if err := t.tx.Commit(); err != nil {
		return t.rollback()
	}

	return nil
}

func (t *Transaction) rollback() error {
	return t.tx.Rollback()
}

// NewTransaction returns a new transaction
func NewTransaction() *Transaction {
	t := &Transaction{}
	t.db, _ = sql.Open("mysql", "root@tcp(localhost:3306)/golang")
	t.Begin()
	return t
}

// NewAccount returns a new account
func NewAccount() *Account {
	a := &Account{}
	a.db, _ = sql.Open("mysql", "root@tcp(localhost:3306)/golang")
	result, _ := a.db.Exec("INSERT INTO account SET balance = ? ", 0)
	id, _ := result.LastInsertId()

	a.ID = id

	return a
}
