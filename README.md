# rinha-2024-crebito-go

## Run

```bash
docker-compose up -d
PORT=9999 DATABASE_URL_CONN=postgres://crebito_user:crebito_pass@localhost:5432/crebito_db go run src/main.go
```

## Run with VSCode

```json
{
    "name": "Launch Package",
    "type": "go",
    "request": "launch",
    "mode": "auto",
    "program": "${workspaceFolder}/src/main.go",
    "env": {
        "PORT": "9999",
        "DATABASE_URL_CONN": "postgres://crebito_user:crebito_pass@localhost:5432/crebito_db"
    }
}
```

## Reset data after execution

```sql
delete from "transaction" 
update client set balance=0
ALTER SEQUENCE transaction_id_seq RESTART WITH 1;
```

## Build and Push Docker image

```bash
docker login
make docker
```
