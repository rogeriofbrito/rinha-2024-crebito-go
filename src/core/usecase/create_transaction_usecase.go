package usecase

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v4"
	"github.com/rogeriofbrito/rinha-2024-crebito-go/src/core/domain"
	external_repository "github.com/rogeriofbrito/rinha-2024-crebito-go/src/core/external/repository"
	"github.com/sarulabs/di"
)

type CreateTransactionUseCase struct {
	Cr external_repository.IClientRepository
	Tr external_repository.ITransactionRepository
}

func (ctuc CreateTransactionUseCase) Execute(dic di.Container, transaction domain.TransactionDomain) (domain.ClientDomain, domain.TransactionDomain, error) {
	tx := dic.Get("tx").(pgx.Tx)

	client, err := ctuc.Cr.GetById(dic, transaction.ClientId)
	if err != nil {
		tx.Rollback(context.Background())
		return domain.ClientDomain{}, domain.TransactionDomain{}, err
	}

	client, err = ctuc.updateClient(dic, client, transaction)
	if err != nil {
		tx.Rollback(context.Background())
		return domain.ClientDomain{}, domain.TransactionDomain{}, err
	}

	transaction, err = ctuc.Tr.Save(dic, transaction)
	if err != nil {
		tx.Rollback(context.Background())
		return domain.ClientDomain{}, domain.TransactionDomain{}, err
	}

	tx.Commit(context.Background())
	return client, transaction, nil
}

func (ctuc CreateTransactionUseCase) updateClient(dic di.Container, client domain.ClientDomain, transaction domain.TransactionDomain) (domain.ClientDomain, error) {
	isValidTransaction, err := ctuc.isValidTransaction(client, transaction)
	if err != nil {
		return domain.ClientDomain{}, err
	}

	if !isValidTransaction {
		return domain.ClientDomain{}, errors.New("invalid transaction") // TODO: create specfic error
	}

	if transaction.Type == domain.Debit {
		client.Balance -= transaction.Value
		return ctuc.Cr.Update(dic, client)
	}

	if transaction.Type == domain.Credit {
		client.Balance += transaction.Value
		return ctuc.Cr.Update(dic, client)
	}

	return domain.ClientDomain{}, fmt.Errorf("invalid transaction type: %d", transaction.Type)
}

func (ctuc CreateTransactionUseCase) isValidTransaction(client domain.ClientDomain, transaction domain.TransactionDomain) (bool, error) {
	switch transaction.Type {
	case domain.Credit:
		return true, nil
	case domain.Debit:
		if client.Balance-transaction.Value+client.Limit < 0 {
			return false, nil
		}
		return true, nil
	}

	return false, fmt.Errorf("invalid transaction type: %d", transaction.Type)
}
