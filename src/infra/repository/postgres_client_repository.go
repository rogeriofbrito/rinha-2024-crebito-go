package infra_repository

import (
	"context"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/rogeriofbrito/rinha-2024-crebito-go/src/core/domain"
)

type PostgresClientRepository struct {
	Conn *pgxpool.Conn
}

func (pcr PostgresClientRepository) GetById(id int64) (domain.ClientDomain, error) {
	pcr.Conn.Ping(context.Background())
	return domain.ClientDomain{}, nil
}
