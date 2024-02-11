package infra_repository

import (
	"context"

	"github.com/jackc/pgx/v4/pgxpool"
)

type PostgresConnUtils struct {
	Conn *pgxpool.Conn
}

func (pcu PostgresConnUtils) InitTransaction() {
	pcu.Conn.Ping(context.Background())
}

func (pcu PostgresConnUtils) CommitTransaction() {
	pcu.Conn.Ping(context.Background())
}

func (pcu PostgresConnUtils) RollbackTransaction() {
	pcu.Conn.Ping(context.Background())
}

func (pcu PostgresConnUtils) CloseConn() {
	pcu.Conn.Ping(context.Background())
}
