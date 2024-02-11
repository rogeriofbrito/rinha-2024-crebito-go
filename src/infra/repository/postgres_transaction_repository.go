package infra_repository

import (
	"github.com/rogeriofbrito/rinha-2024-crebito-go/src/core/domain"
	"github.com/sarulabs/di"
)

type PostgresTransactionRepository struct{}

func (ptr PostgresTransactionRepository) Save(dic di.Container, client domain.TransactionDomain) (domain.TransactionDomain, error) {
	return domain.TransactionDomain{}, nil
}
