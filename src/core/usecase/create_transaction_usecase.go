package usecase

import (
	"github.com/rogeriofbrito/rinha-2024-crebito-go/src/core/domain"
	external_repository "github.com/rogeriofbrito/rinha-2024-crebito-go/src/core/external/repository"
)

type CreateTransactionUseCase struct {
	Cu external_repository.IConnUtils
	Cr external_repository.IClientRepository
	Tr external_repository.ITransactionRepository
}

func (ctuc CreateTransactionUseCase) Execute(transaction domain.TransactionDomain) (domain.TransactionDomain, error) {
	defer ctuc.Cu.CloseConn()

	ctuc.Cu.InitTransaction()

	client, err := ctuc.Cr.GetById(transaction.ClientId)
	if err != nil {
		ctuc.Cu.RollbackTransaction()
		return domain.TransactionDomain{}, err
	}

	doStuf(client)

	transaction, err = ctuc.Tr.Save(transaction)
	if err != nil {
		ctuc.Cu.RollbackTransaction()
		return domain.TransactionDomain{}, err
	}

	ctuc.Cu.CommitTransaction()
	return transaction, nil
}

func doStuf(client domain.ClientDomain) {}
