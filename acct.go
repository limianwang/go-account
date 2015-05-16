package acct

type account struct {
	balance float64
}

func (a *account) Balance() float64 {
	return a.balance
}

type transaction struct {
}

func (t *transaction) MoveMoney(a float64, from *account, to *account) {
}

func (t *transaction) Close() {
}

func NewAccount() *account {
	return &account{}
}

func NewTransaction() *transaction {
	return &transaction{}
}
