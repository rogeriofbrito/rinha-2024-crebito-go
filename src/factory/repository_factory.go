package factory

import (
	"context"
	"os"
	"sync"

	"github.com/jackc/pgx/v4/pgxpool"
	infra_repository "github.com/rogeriofbrito/rinha-2024-crebito-go/src/infra/repository"
)

var postgresPoolLock = &sync.Mutex{}
var singlePostgresPool *pgxpool.Pool

func newPostgresConnPool() *pgxpool.Pool {
	if singlePostgresPool == nil {
		postgresPoolLock.Lock()
		defer postgresPoolLock.Unlock()
		if singlePostgresPool == nil {
			config, err := pgxpool.ParseConfig(os.Getenv("DATABASE_URL_CONN"))
			if err != nil {
				panic(err)
			}

			singlePostgresPool, err = pgxpool.ConnectConfig(context.Background(), config)
			if err != nil {
				panic(err)
			}
		}
	}

	return singlePostgresPool
}

func newPostgresConnUtils(conn *pgxpool.Conn) infra_repository.PostgresConnUtils {
	return infra_repository.PostgresConnUtils{
		Conn: conn,
	}
}

func newPostgresClientRepository(conn *pgxpool.Conn) infra_repository.PostgresClientRepository {
	return infra_repository.PostgresClientRepository{
		Conn: conn,
	}
}

func newPostgresTransactionRepository(conn *pgxpool.Conn) infra_repository.PostgresTransactionRepository {
	return infra_repository.PostgresTransactionRepository{
		Conn: conn,
	}
}
