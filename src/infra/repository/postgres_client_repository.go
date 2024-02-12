package infra_repository

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v4"
	"github.com/rogeriofbrito/rinha-2024-crebito-go/src/core/domain"
	"github.com/sarulabs/di"
)

type PostgresClientRepository struct{}

func (pcr PostgresClientRepository) GetById(dic di.Container, id int64) (domain.ClientDomain, error) {
	tx := dic.Get("tx").(pgx.Tx)

	rows, err := tx.Query(context.Background(), "select id, \"limit\", balance from client where id=$1 FOR UPDATE", id)
	if err != nil {
		return domain.ClientDomain{}, err
	}
	defer rows.Close()

	if !rows.Next() {
		return domain.ClientDomain{}, errors.New("not found")
	}

	client := domain.ClientDomain{}
	err = rows.Scan(&client.Id, &client.Limit, &client.Balance)
	if err != nil {
		return domain.ClientDomain{}, err
	}

	return client, nil
}
