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

func (tc TransactionController) CreateTransaction(dic di.Container, clientId int64, req controller_model.CreateTransactionRequestModel) (controller_model.CreateTransactionResponseModel, error) {
	transactionType, err := getTransactionType(req.Type)
	if err != nil {
		return controller_model.CreateTransactionResponseModel{}, err
	}

	td := domain.TransactionDomain{
		ClientId:    clientId,
		Type:        transactionType,
		Value:       req.Value,
		Description: req.Description,
	}

	cd, _, err := tc.Ctuc.Execute(dic, td)
	if err != nil {
		return controller_model.CreateTransactionResponseModel{}, err
	}

	return controller_model.CreateTransactionResponseModel{
		Limit:   cd.Limit,
		Balance: cd.Balance,
	}, nil
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
