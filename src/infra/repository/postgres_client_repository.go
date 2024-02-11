package infra_repository

import (
	"github.com/rogeriofbrito/rinha-2024-crebito-go/src/core/domain"
)

type PostgresClientRepository struct{}

func (pcr PostgresClientRepository) GetById(id int64) (domain.ClientDomain, error) {
	return domain.ClientDomain{}, nil
}
