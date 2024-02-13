package domain

import "time"

type StatementBalanceDomain struct {
	Total         int64
	StatementDate time.Time
	Limit         int64
}

type StatementTransactionDomain struct {
	Value        int64
	Type         TransactionType
	Description  string
	CarriedOutIn time.Time
}

type StatementDomain struct {
	Balance          StatementBalanceDomain
	LastTransactions []StatementTransactionDomain
}
