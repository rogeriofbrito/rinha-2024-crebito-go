package usecase

import (
	"context"

	"github.com/jackc/pgx/v4"
	"github.com/rogeriofbrito/rinha-2024-crebito-go/src/core/domain"
	external_repository "github.com/rogeriofbrito/rinha-2024-crebito-go/src/core/external/repository"
	"github.com/sarulabs/di"
)

type CreateTransactionUseCase struct {
	Cr external_repository.IClientRepository
	Tr external_repository.ITransactionRepository
}

func (ctuc CreateTransactionUseCase) Execute(dic di.Container, transaction domain.TransactionDomain) (domain.TransactionDomain, error) {
	tx := dic.Get("tx").(pgx.Tx)

	client, err := ctuc.Cr.GetById(dic, transaction.ClientId)
	if err != nil {
		tx.Rollback(context.Background())
		return domain.TransactionDomain{}, err
	}

	doStuf(client)

	transaction, err = ctuc.Tr.Save(dic, transaction)
	if err != nil {
		tx.Rollback(context.Background())
		return domain.TransactionDomain{}, err
	}

	tx.Commit(context.Background())
	return transaction, nil
}

func doStuf(client domain.ClientDomain) {}
