package controller_model

import (
	"github.com/rogeriofbrito/rinha-2024-crebito-go/src/utils"
)

type CreateTransactionRequestModel struct {
	Value       int64  `json:"valor" validate:"required"`
	Type        string `json:"tipo" validate:"oneof=c d,required"`
	Description string `json:"descricao" validate:"min=1,max=10,required"`
}

type CreateTransactionResponseModel struct {
	Limit   int64 `json:"limite"`
	Balance int64 `json:"saldo"`
}

type GetStatementBalanceResponseModel struct {
	Total         int64            `json:"total"`
	StatementDate utils.CustomTime `json:"data_extrato"`
	Limit         int64            `json:"limite"`
}

type GetStatementTransactionResponseModel struct {
	Value        int64            `json:"valor"`
	Type         string           `json:"tipo"`
	Description  string           `json:"descricao"`
	CarriedOutIn utils.CustomTime `json:"realizada_em"`
}

type GetStatementResponseModel struct {
	Balance          GetStatementBalanceResponseModel       `json:"saldo"`
	LastTransactions []GetStatementTransactionResponseModel `json:"ultimas_transacoes"`
}
