package usecase

import (
	"github.com/rogeriofbrito/rinha-2024-crebito-go/src/core/domain"
	"github.com/rogeriofbrito/rinha-2024-crebito-go/src/core/external/repository"
)

type CreateTransactionUseCase struct {
	Cr repository.IClientRepository
	Tr repository.ITransactionRepository
}

func (ctuc CreateTransactionUseCase) Execute(transaction domain.TransactionDomain) (domain.TransactionDomain, error) {
	return domain.TransactionDomain{}, nil
}
