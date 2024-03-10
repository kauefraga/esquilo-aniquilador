package services

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
	"github.com/kauefraga/esquilo-aniquilador/database"
	"github.com/kauefraga/esquilo-aniquilador/internal/domain"
)

var ErroClienteNaoExiste = errors.New("cliente não existe")
var ErroValorDaTransacao = errors.New("o valor da transação deve ser maior que zero")
var ErroTipoDaTransacao = errors.New("a transação deve ser do tipo 'c' (crédito) ou 'd' (débito)")
var ErroDescricao = errors.New("a descrição deve ter de 1 a 10 caractéres")
var ErroTransacaoDebito = errors.New("a transação do tipo 'd' (débito) nunca pode deixar o saldo do cliente menor que seu limite disponível")

// Service sem repository mesmo, só confia kkkkkk
func CreateTransacao(
	clienteId int,
	transacao *domain.TransacaoRequest,
) (*domain.TransacaoResponse, error) {
	// Verifica campos do request
	if transacao.Valor <= 0 {
		return nil, ErroValorDaTransacao
	}

	if transacao.Tipo != "c" && transacao.Tipo != "d" {
		return nil, ErroTipoDaTransacao
	}

	if len(transacao.Descricao) < 1 || len(transacao.Descricao) > 10 {
		return nil, ErroDescricao
	}

	if database.Conn == nil {
		database.GetConnection()
	}

	tx, err := database.Conn.Begin(context.Background())
	if err != nil {
		return nil, err
	}
	defer tx.Rollback(context.Background())

	// Verifica se o ID é de um cliente existente
	var cliente domain.Cliente
	err = tx.
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

	// Verifica regra de negócio
	// Regra de negócio - Uma transação de débito NUNCA pode deixar o saldo do cliente menor que seu limite disponível
	// https://twitter.com/rkauefraga/status/1757524333629464861
	if transacao.Tipo == "d" && cliente.Saldo-transacao.Valor < -cliente.Limite {
		return nil, ErroTransacaoDebito
	}

	cliente.Saldo -= transacao.Valor

	batch := &pgx.Batch{}

	batch.Queue(
		"UPDATE clientes SET saldo = $1 WHERE id = $2",
		cliente.Saldo,
		cliente.ID,
	)
	batch.Queue(
		"INSERT INTO transacoes (valor, tipo, descricao, cliente_id) VALUES ($1, $2, $3, $4)",
		transacao.Valor,
		transacao.Tipo,
		transacao.Descricao,
		cliente.ID,
	)

	s := tx.SendBatch(
		context.Background(),
		batch,
	)
	if err := s.Close(); err != nil {
		return nil, err
	}
	err = tx.Commit(context.Background())
	if err != nil {
		if err.Error() == "no rows in result set" {
			return nil, ErroClienteNaoExiste
		}
		return nil, err
	}

	return &domain.TransacaoResponse{
		Limite: cliente.Limite,
		Saldo:  cliente.Saldo,
	}, nil
}
