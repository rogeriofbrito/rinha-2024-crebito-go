package external_repository

import (
	"github.com/rogeriofbrito/rinha-2024-crebito-go/src/core/domain"
	"github.com/sarulabs/di"
)

type IClientRepository interface {
	GetById(dic di.Container, id int64, options DBOptions) (domain.ClientDomain, error)
	Update(dic di.Container, client domain.ClientDomain, options DBOptions) (domain.ClientDomain, error)
}
