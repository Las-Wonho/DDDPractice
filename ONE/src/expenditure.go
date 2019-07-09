package main

type expenditure struct {
	transaction transaction
	info        expenditureInformation
}

type expenditureInformation struct {
	date        string
	transaction transaction
	spendType   int
}

func (factory expenditureFactory) createExpenditure() expenditure {
	return expenditure{}
}
