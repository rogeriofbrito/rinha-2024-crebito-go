package di

import (
	"context"
	"os"
	"sync"

	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v4/pgxpool"
	external_repository "github.com/rogeriofbrito/rinha-2024-crebito-go/src/core/external/repository"
	"github.com/rogeriofbrito/rinha-2024-crebito-go/src/core/usecase"
	"github.com/rogeriofbrito/rinha-2024-crebito-go/src/infra/controller"
	controller_adapter "github.com/rogeriofbrito/rinha-2024-crebito-go/src/infra/controller/adapter"
	infra_repository "github.com/rogeriofbrito/rinha-2024-crebito-go/src/infra/repository"
	"github.com/sarulabs/di"
)

func GetDiContainer() di.Container {
	builder, _ := di.NewBuilder()

	builder.Add(di.Def{
		Name:  "fiber-controller-adapter",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
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
			return fiber.New(), nil
		},
	})

	builder.Add(di.Def{
		Name:  "transaction-controller",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			return controller.TransactionController{
				Ctuc: ctn.Get("create-transaction-use-case").(usecase.CreateTransactionUseCase),
			}, nil
		},
	})

	builder.Add(di.Def{
		Name:  "create-transaction-use-case",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
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
			return infra_repository.PostgresClientRepository{}, nil
		},
	})

	builder.Add(di.Def{
		Name:  "transaction-repository",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			return infra_repository.PostgresTransactionRepository{}, nil
		},
	})

	builder.Add(di.Def{
		Name:  "conn",
		Scope: di.Request,
		Build: func(ctn di.Container) (interface{}, error) {
			pool := ctn.Get("pool").(*pgxpool.Pool)
			conn, err := pool.Acquire(context.Background())
			if err != nil {
				return nil, err
			}

			return conn, nil
		},
		Close: func(obj interface{}) error {
			obj.(*pgxpool.Conn).Release()
			return nil
		},
	})

	builder.Add(di.Def{
		Name:  "pool",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			config, err := pgxpool.ParseConfig(os.Getenv("DATABASE_URL_CONN"))
			if err != nil {
				return nil, err
			}

			pool, err := pgxpool.ConnectConfig(context.Background(), config)
			if err != nil {
				return nil, err
			}

			return pool, nil
		},
		Close: func(obj interface{}) error {
			obj.(*pgxpool.Pool).Close()
			return nil
		},
	})

	return builder.Build()
}

var diContainer di.Container
var lock = &sync.Mutex{}

/*
func GetDiContainer() di.Container {
	if diContainer == nil {
		lock.Lock()
		defer lock.Unlock()
		if diContainer == nil {
			diContainer = getDiContainer()
		}
	}

	return diContainer
}
*/
