package usecase

import (
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/rogeriofbrito/rinha-2024-crebito-go/src/core/domain"
	external_repository "github.com/rogeriofbrito/rinha-2024-crebito-go/src/core/external/repository"
	"github.com/sarulabs/di"
	log "github.com/sirupsen/logrus"
)

type CreateTransactionUseCase struct {
	Cr external_repository.IClientRepository
	Tr external_repository.ITransactionRepository
}

func (ctuc CreateTransactionUseCase) Execute(dic di.Container, transaction domain.TransactionDomain) (domain.TransactionDomain, error) {
	conn := dic.Get("conn").(*pgxpool.Conn)
	log.WithFields(log.Fields{
		"pid": conn.Conn().PgConn().PID(),
	}).Debug("Connection acquired from pool")

	client, err := ctuc.Cr.GetById(dic, transaction.ClientId)
	if err != nil {
		return domain.TransactionDomain{}, err
	}

	doStuf(client)

	transaction, err = ctuc.Tr.Save(dic, transaction)
	if err != nil {
		return domain.TransactionDomain{}, err
	}

	return transaction, nil
}

func doStuf(client domain.ClientDomain) {}
