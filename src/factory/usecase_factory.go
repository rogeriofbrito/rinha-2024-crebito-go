package factory

import (
	"context"

	"github.com/rogeriofbrito/rinha-2024-crebito-go/src/core/usecase"
)

func newCreateTransactionUseCase() usecase.CreateTransactionUseCase {
	pool := newPostgresConnPool()

	conn, err := pool.Acquire(context.Background())
	if err != nil {
		panic(err)
	}

	return usecase.CreateTransactionUseCase{
		Cu: newPostgresConnUtils(conn),             // for now, only Postgers implemented
		Cr: newPostgresClientRepository(conn),      // for now, only Postgers implemented
		Tr: newPostgresTransactionRepository(conn), // for now, only Postgers implemented
	}
}
