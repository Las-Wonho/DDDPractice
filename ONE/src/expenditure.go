package main

import "time"

type expenditure struct {
	transaction transaction
	info        expenditureInformation
}

type expenditureInformation struct {
	date      string
	spendType int
	reason    string
	who       string
}

type expenditureFactory struct{}

func (factory expenditureFactory) createExpenditure(spendType int, spendedMoney int) expenditure {
	if spendType < lending || spendType > otherThing {
		spendType = otherThing
	}
	if spendedMoney < 0 {
		spendedMoney := -spendedMoney
	}
	info := expenditureInformation{date: time.ANSIC, spendType: spendType}
	ts := transaction{moneyChange: spendType}
	return expenditure{info: info, transaction: ts}
}

const (
	lending int = 1 + iota
	meal
	snack
	transportation
	subscription
	lisence
	otherThing
)
