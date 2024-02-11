package infra_repository

import (
	"context"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/rogeriofbrito/rinha-2024-crebito-go/src/core/domain"
)

type PostgresTransactionRepository struct {
	Conn *pgxpool.Conn
}

func (ptr PostgresTransactionRepository) Save(client domain.TransactionDomain) (domain.TransactionDomain, error) {
	ptr.Conn.Ping(context.Background())
	return domain.TransactionDomain{}, nil
}
