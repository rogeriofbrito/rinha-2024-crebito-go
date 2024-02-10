package domain

type TransactionType uint8

const (
	Credit TransactionType = iota
	Debit
)

type TransactionDomain struct {
	Id          int64
	Type        TransactionType
	Value       int64
	Description string
}
