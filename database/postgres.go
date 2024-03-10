package database

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

var Conn *pgxpool.Pool

// POR FAVOR, FECHA A CONEXÃO DEPOIS DE USAR OU METE `defer db.Close()`
func GetConnection() {
	connString := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)

	config, err := pgxpool.ParseConfig(connString)
	if err != nil {
		panic(err)
	}

	db, err := pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		panic(err)
	}

	err = db.Ping(context.Background())
	if err != nil {
		panic(err)
	}

	if Conn == nil {
		Conn = db
	}
}
