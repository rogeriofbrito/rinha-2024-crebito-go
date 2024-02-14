package usecase

import (
	"context"
	"errors"
	"fmt"

	"github.com/gofiber/fiber/v2/log"
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

	client, err := ctuc.Cr.GetById(dic, transaction.ClientId, external_repository.DBOptions{
		LockMode: external_repository.Pessimistic,
	})
	if err != nil {
		tx.Rollback(context.Background())
		return domain.ClientDomain{}, domain.TransactionDomain{}, err
	}

	client, err = ctuc.updateClient(dic, client, transaction)
	if err != nil {
		tx.Rollback(context.Background())
		return domain.ClientDomain{}, domain.TransactionDomain{}, err
	}

	transaction, err = ctuc.Tr.Save(dic, transaction, external_repository.DBOptions{})
	if err != nil {
		tx.Rollback(context.Background())
		return domain.ClientDomain{}, domain.TransactionDomain{}, err
	}

	tx.Commit(context.Background())
	return client, transaction, nil
}

func (ctuc CreateTransactionUseCase) updateClient(dic di.Container, client domain.ClientDomain, transaction domain.TransactionDomain) (domain.ClientDomain, error) {
	isValidTransaction, msg := ctuc.isValidTransaction(client, transaction)

	if !isValidTransaction {
		log.Error(msg)
		return domain.ClientDomain{}, errors.New("invalid transaction") // TODO: create specfic error
	}

	if transaction.Type == domain.Debit {
		client.Balance -= transaction.Value
		return ctuc.Cr.Update(dic, client, external_repository.DBOptions{})
	}

	if transaction.Type == domain.Credit {
		client.Balance += transaction.Value
		return ctuc.Cr.Update(dic, client, external_repository.DBOptions{})
	}

	return domain.ClientDomain{}, fmt.Errorf("invalid transaction type: %d", transaction.Type)
}

func (ctuc CreateTransactionUseCase) isValidTransaction(client domain.ClientDomain, transaction domain.TransactionDomain) (bool, string) {
	switch transaction.Type {
	case domain.Credit:
		return true, "valid transaction - credit"
	case domain.Debit:
		if client.Balance-transaction.Value+client.Limit < 0 {
			return false, "invalid transaction - insuficient limit"
		}
		return true, "valid transaction - debit"
	}

	errMsg := fmt.Sprintf("invalid transaction type: %d", transaction.Type)
	return false, errMsg
}
