package controller

import (
	"github.com/rogeriofbrito/rinha-2024-crebito-go/src/core/domain"
	"github.com/rogeriofbrito/rinha-2024-crebito-go/src/core/usecase"
	controller_model "github.com/rogeriofbrito/rinha-2024-crebito-go/src/infra/controller/model"
	"github.com/sarulabs/di"
)

type TransactionController struct {
	Ctuc usecase.CreateTransactionUseCase
}

func (tc TransactionController) CreateTransaction(dic di.Container, tm controller_model.TransactionModel) (controller_model.TransactionModel, error) {
	td, err := tc.Ctuc.Execute(dic, domain.TransactionDomain{})
	if err != nil {
		return controller_model.TransactionModel{}, err
	}

	print(td.Id)

	return controller_model.TransactionModel{}, nil
}
