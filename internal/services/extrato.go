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
	// Pega conexão com banco de dados
	conn := database.GetConnection()
	defer conn.Close()

	// Verifica se o ID é de um cliente existente e pega ele
	var cliente domain.Cliente
	err := conn.
		QueryRow(
			context.Background(),
			"SELECT * FROM clientes WHERE id = $1",
			clienteId,
		).
		Scan(
			&cliente.ID,
			&cliente.Nome,
			&cliente.Limite,
			&cliente.Saldo,
		)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErroClienteNaoExiste
		}

		return nil, err
	}

	rows, err := conn.Query(
		context.Background(),
		"SELECT valor, tipo, descricao, realizada_em FROM transacoes WHERE cliente_id = $1 ORDER BY realizada_em DESC LIMIT 10",
		cliente.ID,
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
