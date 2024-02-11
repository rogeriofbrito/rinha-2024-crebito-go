package external_repository

type IConnUtils interface {
	InitTransaction()
	CommitTransaction()
	RollbackTransaction()
	CloseConn()
}
