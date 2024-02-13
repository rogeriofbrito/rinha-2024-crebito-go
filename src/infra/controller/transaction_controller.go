package controller

import (
	"fmt"

	"github.com/rogeriofbrito/rinha-2024-crebito-go/src/core/domain"
	"github.com/rogeriofbrito/rinha-2024-crebito-go/src/core/usecase"
	controller_model "github.com/rogeriofbrito/rinha-2024-crebito-go/src/infra/controller/model"
	"github.com/rogeriofbrito/rinha-2024-crebito-go/src/utils"
	"github.com/sarulabs/di"
)

type TransactionController struct {
	Ctuc usecase.CreateTransactionUseCase
	Gsuc usecase.GetStatementUseCase
}

func (tc TransactionController) CreateTransaction(dic di.Container, clientId int64, req controller_model.CreateTransactionRequestModel) (controller_model.CreateTransactionResponseModel, error) {
	transactionType, err := tc.getTransactionType(req.Type)
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

func (tc TransactionController) GetStatement(dic di.Container, clientId int64) (controller_model.GetStatementResponseModel, error) {
	sd, err := tc.Gsuc.Execute(dic, clientId)
	if err != nil {
		return controller_model.GetStatementResponseModel{}, err
	}

	return controller_model.GetStatementResponseModel{
		Balance: controller_model.GetStatementBalanceResponseModel{
			Total:         sd.Balance.Total,
			StatementDate: utils.CustomTime(sd.Balance.StatementDate),
			Limit:         sd.Balance.Limit,
		},
		LastTransactions: tc.convertTransactions(sd.LastTransactions),
	}, nil
}

func (tc TransactionController) getTransactionType(transactionTypeStr string) (domain.TransactionType, error) {
	if transactionTypeStr == "c" {
		return domain.Credit, nil
	} else if transactionTypeStr == "d" {
		return domain.Debit, nil
	} else {
		return 0, fmt.Errorf("invalid transaction type: %s", transactionTypeStr) // TODO: create specfic error
	}
}

func (tc TransactionController) convertTransactions(stds []domain.StatementTransactionDomain) []controller_model.GetStatementTransactionResponseModel {
	strms := []controller_model.GetStatementTransactionResponseModel{}

	for _, v := range stds {
		strms = append(strms, controller_model.GetStatementTransactionResponseModel{
			Value:        v.Value,
			Type:         v.Type,
			Description:  v.Description,
			CarriedOutIn: utils.CustomTime(v.CarriedOutIn),
		})
	}

	return strms
}
