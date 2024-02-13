package external_repository

type LockType string

const (
	Pessimistic LockType = "PESSIMISTIC"
	Optimistic  LockType = "OPTIMISTIC"
)

type DBOptions struct {
	LockMode LockType
}
