package infra_repository

import (
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/rogeriofbrito/rinha-2024-crebito-go/src/core/domain"
	"github.com/sarulabs/di"
	log "github.com/sirupsen/logrus"
)

type PostgresTransactionRepository struct{}

func (ptr PostgresTransactionRepository) Save(dic di.Container, client domain.TransactionDomain) (domain.TransactionDomain, error) {
	conn := dic.Get("conn").(*pgxpool.Conn)
	log.WithFields(log.Fields{
		"pid": conn.Conn().PgConn().PID(),
	}).Debug("Connection acquired from pool")

	return domain.TransactionDomain{}, nil
}
