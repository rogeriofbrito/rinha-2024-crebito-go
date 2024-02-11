package external_repository

import "github.com/rogeriofbrito/rinha-2024-crebito-go/src/core/domain"

type IClientRepository interface {
	GetById(id int64) (domain.ClientDomain, error)
}
