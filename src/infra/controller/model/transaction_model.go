package controller_model

type CreateTransactionRequestModel struct {
	Value       int64  `json:"valor"`
	Type        string `json:"tipo"`
	Description string `json:"descricao"`
}

type CreateTransactionResponseModel struct {
	Limit   int64 `json:"limite"`
	Balance int64 `json:"saldo"`
}
