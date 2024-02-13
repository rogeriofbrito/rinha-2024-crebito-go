package controller_model

import (
	"time"

	"github.com/rogeriofbrito/rinha-2024-crebito-go/src/core/domain"
)

type CreateTransactionRequestModel struct {
	Value       int64  `json:"valor"`
	Type        string `json:"tipo"`
	Description string `json:"descricao"`
}

type CreateTransactionResponseModel struct {
	Limit   int64 `json:"limite"`
	Balance int64 `json:"saldo"`
}

type GetStatementBalanceResponseModel struct {
	Total         int64     `json:"total"`
	StatementDate time.Time `json:"data_extrato"`
	Limit         int64     `json:"limite"`
}

type GetStatementTransactionResponseModel []struct {
	Value        int64                  `json:"valor"`
	Type         domain.TransactionType `json:"tipo"`
	Description  string                 `json:"descricao"`
	CarriedOutIn time.Time              `json:"realizada_em"`
}

type GetStatementResponseModel struct {
	Balance          GetStatementBalanceResponseModel     `json:"saldo"`
	LastTransactions GetStatementTransactionResponseModel `json:"ultimas_transacoes"`
}
