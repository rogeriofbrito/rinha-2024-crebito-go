package usecase

import (
	"context"
	"time"

	"github.com/jackc/pgx/v4"
	"github.com/rogeriofbrito/rinha-2024-crebito-go/src/core/domain"
	external_repository "github.com/rogeriofbrito/rinha-2024-crebito-go/src/core/external/repository"
	"github.com/sarulabs/di"
)

type GetStatementUseCase struct {
	Cr external_repository.IClientRepository
	Tr external_repository.ITransactionRepository
}

func (gsuc GetStatementUseCase) Execute(dic di.Container, clientId int64) (domain.StatementDomain, error) {
	tx := dic.Get("tx").(pgx.Tx) // TODO: dont open connection ?

	client, err := gsuc.Cr.GetById(dic, clientId, external_repository.DBOptions{})
	if err != nil {
		tx.Rollback(context.Background())
		return domain.StatementDomain{}, err
	}

	transactions, err := gsuc.Tr.GetByClientIdLimit(dic, clientId, 10, external_repository.DBOptions{})
	if err != nil {
		tx.Rollback(context.Background())
		return domain.StatementDomain{}, err
	}

	sd := domain.StatementDomain{
		Balance: domain.StatementBalanceDomain{
			Total:         client.Balance,
			StatementDate: time.Now(),
			Limit:         client.Limit,
		},
		LastTransactions: gsuc.convertTransactions(transactions),
	}

	tx.Commit(context.Background())
	return sd, nil
}

func (gsuc GetStatementUseCase) convertTransactions(tds []domain.TransactionDomain) []domain.StatementTransactionDomain {
	stds := []domain.StatementTransactionDomain{}

	for _, v := range tds {
		stds = append(stds, domain.StatementTransactionDomain{
			Value:        v.Value,
			Type:         v.Type,
			Description:  v.Description,
			CarriedOutIn: v.CreatedAt,
		})
	}

	return stds
}
