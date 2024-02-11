package controller

import (
	"github.com/rogeriofbrito/rinha-2024-crebito-go/src/core/domain"
	"github.com/rogeriofbrito/rinha-2024-crebito-go/src/core/usecase"
	controller_model "github.com/rogeriofbrito/rinha-2024-crebito-go/src/infra/controller/model"
)

type TransactionController struct {
	Ctuc usecase.CreateTransactionUseCase
}

func (tc TransactionController) CreateTransaction(tm controller_model.TransactionModel) (controller_model.TransactionModel, error) {
	td, err := tc.Ctuc.Execute(domain.TransactionDomain{})
	if err != nil {
		return controller_model.TransactionModel{}, err
	}

	print(td.Id)

	return controller_model.TransactionModel{}, nil
}
