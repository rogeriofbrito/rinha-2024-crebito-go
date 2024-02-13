package usecase

import (
	"time"

	"github.com/rogeriofbrito/rinha-2024-crebito-go/src/core/domain"
	external_repository "github.com/rogeriofbrito/rinha-2024-crebito-go/src/core/external/repository"
	"github.com/sarulabs/di"
)

type GetStatementUseCase struct {
	Cr external_repository.IClientRepository
	Tr external_repository.ITransactionRepository
}

func (gsuc GetStatementUseCase) Execute(dic di.Container, clientId int64) (domain.StatementDomain, error) {
	client, err := gsuc.Cr.GetById(dic, clientId, external_repository.DBOptions{})
	if err != nil {
		return domain.StatementDomain{}, err
	}

	return domain.StatementDomain{
		Balance: domain.StatementBalanceDomain{
			Total:         client.Balance,
			StatementDate: time.Now(),
			Limit:         client.Limit,
		},
		LastTransactions: []domain.StatementTransactionDomain{},
	}, nil
}
