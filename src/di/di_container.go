package di

import (
	"context"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v4/pgxpool"
	external_repository "github.com/rogeriofbrito/rinha-2024-crebito-go/src/core/external/repository"
	"github.com/rogeriofbrito/rinha-2024-crebito-go/src/core/usecase"
	"github.com/rogeriofbrito/rinha-2024-crebito-go/src/infra/controller"
	controller_adapter "github.com/rogeriofbrito/rinha-2024-crebito-go/src/infra/controller/adapter"
	infra_repository "github.com/rogeriofbrito/rinha-2024-crebito-go/src/infra/repository"
	"github.com/sarulabs/di"
	log "github.com/sirupsen/logrus"
)

func GetDiContainer() di.Container {
	builder, _ := di.NewBuilder()

	builder.Add(di.Def{
		Name:  "fiber-controller-adapter",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			log.Debug("Building fiber-controller-adapter...")
			return controller_adapter.FiberControllerAdapter{
				Tc:  ctn.Get("transaction-controller").(controller.TransactionController),
				App: ctn.Get("fiber").(*fiber.App),
			}, nil
		},
	})

	builder.Add(di.Def{
		Name:  "fiber",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			log.Debug("Building fiber...")
			return fiber.New(fiber.Config{
				DisableStartupMessage: true,
			}), nil
		},
	})

	builder.Add(di.Def{
		Name:  "transaction-controller",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			log.Debug("Building transaction-controller...")
			return controller.TransactionController{
				Ctuc: ctn.Get("create-transaction-use-case").(usecase.CreateTransactionUseCase),
			}, nil
		},
	})

	builder.Add(di.Def{
		Name:  "create-transaction-use-case",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			log.Debug("Building create-transaction-use-case...")
			return usecase.CreateTransactionUseCase{
				Cr: ctn.Get("client-repository").(external_repository.IClientRepository),
				Tr: ctn.Get("transaction-repository").(external_repository.ITransactionRepository),
			}, nil
		},
	})

	builder.Add(di.Def{
		Name:  "client-repository",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			log.Debug("Building client-repository...")
			return infra_repository.PostgresClientRepository{}, nil
		},
	})

	builder.Add(di.Def{
		Name:  "transaction-repository",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			log.Debug("Building transaction-repository...")
			return infra_repository.PostgresTransactionRepository{}, nil
		},
	})

	builder.Add(di.Def{
		Name:  "conn",
		Scope: di.Request,
		Build: func(ctn di.Container) (interface{}, error) {
			log.Debug("Building conn...")
			pool := ctn.Get("pool").(*pgxpool.Pool)
			conn, err := pool.Acquire(context.Background())
			if err != nil {
				return nil, err
			}

			log.WithFields(log.Fields{
				"pid": conn.Conn().PgConn().PID(),
			}).Debug("Connection acquired from pool")

			return conn, nil
		},
		Close: func(obj interface{}) error {
			conn := obj.(*pgxpool.Conn)
			log.WithFields(log.Fields{
				"pid": conn.Conn().PgConn().PID(),
			}).Debug("Closing conn...")
			conn.Release()
			return nil
		},
	})

	builder.Add(di.Def{
		Name:  "pool",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			log.Debug("Building pool...")
			config, err := pgxpool.ParseConfig(os.Getenv("DATABASE_URL_CONN"))
			if err != nil {
				return nil, err
			}

			pool, err := pgxpool.ConnectConfig(context.Background(), config)
			if err != nil {
				return nil, err
			}

			log.WithFields(log.Fields{
				"MaxConnLifetime":       pool.Config().MaxConnLifetime,
				"MaxConnLifetimeJitter": pool.Config().MaxConnLifetimeJitter,
				"MaxConnIdleTime":       pool.Config().MaxConnIdleTime,
				"MaxConns":              pool.Config().MaxConns,
				"MinConns":              pool.Config().MinConns,
				"HealthCheckPeriod":     pool.Config().HealthCheckPeriod,
				"LazyConnect":           pool.Config().LazyConnect,
			}).Debug("Pool created")

			return pool, nil
		},
		Close: func(obj interface{}) error {
			log.Debug("Closing pool...")
			obj.(*pgxpool.Pool).Close()
			return nil
		},
	})

	return builder.Build()
}
