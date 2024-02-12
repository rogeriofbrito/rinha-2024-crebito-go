package infra_repository

import (
	"context"
	"errors"
	"time"

	"github.com/jackc/pgx/v4"
	"github.com/rogeriofbrito/rinha-2024-crebito-go/src/core/domain"
	"github.com/sarulabs/di"
)

type PostgresTransactionRepository struct{}

func (ptr PostgresTransactionRepository) Save(dic di.Container, transaction domain.TransactionDomain) (domain.TransactionDomain, error) {
	tx := dic.Get("tx").(pgx.Tx)

	transaction.CreatedAt = time.Now()

	rows, err := tx.Query(context.Background(), "insert into transaction (client_id, type, value, description, created_at) VALUES ($1, $2, $3, $4, $5) RETURNING id",
		transaction.ClientId, transaction.Type, transaction.Value, transaction.Description, time.Now())
	if err != nil {
		return domain.TransactionDomain{}, err
	}
	defer rows.Close()

	if !rows.Next() {
		return domain.TransactionDomain{}, errors.New("not found")
	}

	err = rows.Scan(&transaction.Id)
	if err != nil {
		return domain.TransactionDomain{}, err
	}

	return transaction, nil
}
