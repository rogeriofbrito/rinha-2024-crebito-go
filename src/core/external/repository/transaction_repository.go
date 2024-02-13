package external_repository

import (
	"github.com/rogeriofbrito/rinha-2024-crebito-go/src/core/domain"
	"github.com/sarulabs/di"
)

type ITransactionRepository interface {
	GetByClientIdLimit(dic di.Container, clientId int64, limit int64, options DBOptions) ([]domain.TransactionDomain, error)
	Save(dic di.Container, transaction domain.TransactionDomain, options DBOptions) (domain.TransactionDomain, error)
}
