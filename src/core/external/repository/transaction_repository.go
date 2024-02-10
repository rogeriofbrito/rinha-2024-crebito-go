package repository

import "github.com/rogeriofbrito/rinha-2024-crebito-go/src/core/domain"

type ITransactionRepository interface {
	Save(client domain.TransactionDomain) (domain.TransactionDomain, error)
}
