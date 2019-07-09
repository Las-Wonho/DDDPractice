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

func (factory expenditureFactory) createExpenditure(spendType int) expenditure {
	if spendType < lending || spendType > otherThing {
		spendType = otherThing
	}
	info := expenditureInformation{date: time.ANSIC, spendType: spendType}
	return expenditure{info: info}
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
