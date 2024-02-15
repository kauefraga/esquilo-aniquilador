package database

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

// POR FAVOR, FECHA A CONEXÃO DEPOIS DE USAR OU METE `defer db.Close()`
func GetConnection() *pgxpool.Pool {
	dbpool, err := pgxpool.New(
		context.Background(),
		"postgres://admin:rinha@localhost:5432/esquilo-aniquilador",
	)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to create connection pool: %v\n", err)
		os.Exit(1)
	}

	return dbpool
}
