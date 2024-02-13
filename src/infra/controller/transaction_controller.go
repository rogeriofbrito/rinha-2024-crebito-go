package controller

import (
	"fmt"

	"github.com/rogeriofbrito/rinha-2024-crebito-go/src/core/domain"
	"github.com/rogeriofbrito/rinha-2024-crebito-go/src/core/usecase"
	controller_model "github.com/rogeriofbrito/rinha-2024-crebito-go/src/infra/controller/model"
	"github.com/sarulabs/di"
)

type TransactionController struct {
	Ctuc usecase.CreateTransactionUseCase
}

func (tc TransactionController) CreateTransaction(dic di.Container, clientId int64, tm controller_model.TransactionModel) (controller_model.TransactionModel, error) {
	transactionType, err := getTransactionType(tm.Type)
	if err != nil {
		return controller_model.TransactionModel{}, err
	}

	td := domain.TransactionDomain{
		ClientId:    clientId,
		Type:        transactionType,
		Value:       tm.Value,
		Description: tm.Description,
	}

	_, err = tc.Ctuc.Execute(dic, td)
	if err != nil {
		return controller_model.TransactionModel{}, err
	}

	return controller_model.TransactionModel{}, nil
}

func getTransactionType(transactionTypeStr string) (domain.TransactionType, error) {
	if transactionTypeStr == "c" {
		return domain.Credit, nil
	} else if transactionTypeStr == "d" {
		return domain.Debit, nil
	} else {
		return 0, fmt.Errorf("invalid transaction type: %s", transactionTypeStr) // TODO: create specfic error
	}
}
