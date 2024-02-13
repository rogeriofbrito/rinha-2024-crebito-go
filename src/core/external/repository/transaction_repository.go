package external_repository

import (
	"github.com/rogeriofbrito/rinha-2024-crebito-go/src/core/domain"
	"github.com/sarulabs/di"
)

type ITransactionRepository interface {
	Save(dic di.Container, transaction domain.TransactionDomain, options DBOptions) (domain.TransactionDomain, error)
}
