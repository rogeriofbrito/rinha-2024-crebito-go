package domain

import "time"

type TransactionType uint8

const (
	Credit TransactionType = iota
	Debit
)

type TransactionDomain struct {
	Id          int64
	ClientId    int64
	Type        TransactionType
	Value       int64
	Description string
	CreatedAt   time.Time
}
