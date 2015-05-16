package acct

import "testing"

func TestBasicTransaction(t *testing.T) {
	useracct := NewAccount()
	revenueacct := NewAccount()
	taxacct := NewAccount()
	commissionacct := NewAccount()

	retailprice := 10.00
	taxrate := .07
	commissionrate := .1
	taxamt := retailprice * taxrate
	price := retailprice + taxamt
	commissionamt := retailprice * commissionrate

	// Prepare a transaction. As a transaction, the movements should all be
	// committed together or rolled back together, depending on whether there
	// were any errors.
	trans := NewTransaction()
	trans.MoveMoney(price, useracct, revenueacct)
	trans.MoveMoney(taxamt, revenueacct, taxacct)
	trans.MoveMoney(commissionamt, revenueacct, commissionacct)

	shouldbe := 0.0
	was := useracct.Balance()
	if was != shouldbe {
		t.Fatalf("Money was moved before the transaction was closed.")
	}

	// until the transaction is closed/committed, the balances shouldn't change
	trans.Close()

	shouldbe = -10.70
	was = useracct.Balance()
	if was != shouldbe {
		t.Fatalf("Wrong user balance. Was %g, should be %g", was, shouldbe)
	}

	shouldbe = 9.00
	was = revenueacct.Balance()
	if was != shouldbe {
		t.Fatalf("Wrong revenue balance. Was %g, should be %g", was, shouldbe)
	}

	shouldbe = .7
	was = taxacct.Balance()
	if was != shouldbe {
		t.Fatalf("Wrong tax balance. Was %g, should be %g", was, shouldbe)
	}

	shouldbe = 1.00
	was = commissionacct.Balance()
	if was != shouldbe {
		t.Fatalf("Wrong commission balance. Was %g, should be %g", was, shouldbe)
	}

}
