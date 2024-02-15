package infra_repository

import (
	"context"
	"errors"
	"time"

	"github.com/jackc/pgx/v4"
	"github.com/rogeriofbrito/rinha-2024-crebito-go/src/core/domain"
	external_repository "github.com/rogeriofbrito/rinha-2024-crebito-go/src/core/external/repository"
	"github.com/sarulabs/di"
)

type PostgresTransactionRepository struct{}

func (ptr PostgresTransactionRepository) GetByClientIdLimit(dic di.Container, clientId int64, limit int64, options external_repository.DBOptions) ([]domain.TransactionDomain, error) {
	tx := dic.Get("tx").(pgx.Tx)

	query := "select id, client_id, type, value, description, created_at from transaction where client_id=$1 order by created_at desc limit $2"
	rows, err := tx.Query(context.Background(), query, clientId, limit)
	if err != nil {
		return []domain.TransactionDomain{}, err
	}
	defer rows.Close()

	transactions := []domain.TransactionDomain{}
	for rows.Next() {
		transaction := domain.TransactionDomain{}
		err = rows.Scan(&transaction.Id, &transaction.ClientId, &transaction.Type, &transaction.Value, &transaction.Description, &transaction.CreatedAt)
		if err != nil {
			return []domain.TransactionDomain{}, err
		}
		transactions = append(transactions, transaction)
	}

	return transactions, nil
}

func (ptr PostgresTransactionRepository) Save(dic di.Container, transaction domain.TransactionDomain, options external_repository.DBOptions) (domain.TransactionDomain, error) {
	tx := dic.Get("tx").(pgx.Tx)

	transaction.CreatedAt = time.Now()

	rows, err := tx.Query(context.Background(), "insert into transaction (client_id, type, value, description, created_at) VALUES ($1, $2, $3, $4, $5) returning id", //TODO: return all values
		transaction.ClientId, transaction.Type, transaction.Value, transaction.Description, time.Now())
	if err != nil {
		return domain.TransactionDomain{}, err
	}
	defer rows.Close()

	if !rows.Next() {
		return domain.TransactionDomain{}, errors.New("not found") // TODO: create specfic error
	}

	err = rows.Scan(&transaction.Id)
	if err != nil {
		return domain.TransactionDomain{}, err
	}

	return transaction, nil
}
