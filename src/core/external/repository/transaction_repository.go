package external_repository

import "github.com/rogeriofbrito/rinha-2024-crebito-go/src/core/domain"

type ITransactionRepository interface {
	Save(transaction domain.TransactionDomain) (domain.TransactionDomain, error)
}
