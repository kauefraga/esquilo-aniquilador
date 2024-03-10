package services

import (
	"context"
	"errors"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/kauefraga/esquilo-aniquilador/database"
	"github.com/kauefraga/esquilo-aniquilador/internal/domain"
)

func GetExtrato(clienteId int) (*domain.Extrato, error) {
	if clienteId < 1 || clienteId > 5 {
		return nil, ErroClienteNaoExiste
	}

	var cliente domain.Cliente

	if database.Conn == nil {
		database.GetConnection()
	}

	err := database.Conn.
		QueryRow(
			context.Background(),
			"SELECT limite, saldo FROM clientes WHERE id = $1",
			clienteId,
		).
		Scan(
			&cliente.Limite,
			&cliente.Saldo,
		)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErroClienteNaoExiste
		}

		return nil, err
	}

	rows, err := database.Conn.Query(
		context.Background(),
		"SELECT valor, tipo, descricao, realizada_em FROM transacoes WHERE cliente_id = $1 ORDER BY realizada_em DESC LIMIT 10",
		clienteId,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	ultimasTransacoes, err := pgx.CollectRows(rows, pgx.RowToStructByName[domain.UltimaTransacao])
	if err != nil {
		return nil, err
	}

	return &domain.Extrato{
		Saldo: domain.Saldo{
			Total:       cliente.Saldo,
			DataExtrato: time.Now(),
			Limite:      cliente.Limite,
		},
		UltimasTransacoes: ultimasTransacoes,
	}, nil
}
