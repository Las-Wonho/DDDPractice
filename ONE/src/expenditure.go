package main

type expenditure struct {
	transaction transaction
	info        expenditureInformation
}

type expenditureInformation struct{}

type expenditureFactory struct {
}
