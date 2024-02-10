package controller

import controller_model "github.com/rogeriofbrito/rinha-2024-crebito-go/src/infra/controller/model"

type TransactionController struct {
}

func (tc TransactionController) CreateTransaction(tm controller_model.TransactionModel) (controller_model.TransactionModel, error) {
	return controller_model.TransactionModel{}, nil
}
