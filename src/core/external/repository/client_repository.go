package external_repository

import (
	"github.com/rogeriofbrito/rinha-2024-crebito-go/src/core/domain"
	"github.com/sarulabs/di"
)

type IClientRepository interface {
	GetById(dic di.Container, id int64) (domain.ClientDomain, error)
}
