package infra_repository

import (
	"github.com/rogeriofbrito/rinha-2024-crebito-go/src/core/domain"
)

type PostgresTransactionRepository struct{}

func (ptr PostgresTransactionRepository) Save(client domain.TransactionDomain) (domain.TransactionDomain, error) {
	return domain.TransactionDomain{}, nil
}
