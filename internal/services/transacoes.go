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
var ErroTransacao = errors.New("a transação do tipo 'd' (débito) nunca pode deixar o saldo do cliente menor que seu limite disponível")

// Service sem repository mesmo, só confia kkkkkk
func CreateTransacao(
	clienteId int,
	transacao *domain.TransacaoRequest,
) (*domain.TransacaoResponse, error) {
	if clienteId < 1 || clienteId > 5 {
		return nil, ErroClienteNaoExiste
	}

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

	var limite int
	var saldo int
	err = tx.
		QueryRow(
			context.Background(),
			"SELECT limite, saldo FROM clientes WHERE id = $1 FOR UPDATE",
			clienteId,
		).
		Scan(
			&limite,
			&saldo,
		)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErroClienteNaoExiste
		}

		return nil, err
	}

	var novo_saldo int

	if transacao.Tipo == "d" {
		novo_saldo = saldo - transacao.Valor
	} else {
		novo_saldo = saldo + transacao.Valor
	}

	// Verifica regra de negócio
	// Regra de negócio - Uma transação de débito NUNCA pode deixar o saldo do cliente menor que seu limite disponível
	// https://twitter.com/rkauefraga/status/1757524333629464861
	if novo_saldo < limite*-1 {
		return nil, ErroTransacao
	}

	_, err = tx.Exec(
		context.Background(),
		"INSERT INTO transacoes (valor, tipo, descricao, cliente_id) VALUES ($1, $2, $3, $4)",
		transacao.Valor,
		transacao.Tipo,
		transacao.Descricao,
		clienteId,
	)
	if err != nil {
		return nil, err
	}

	_, err = tx.Exec(
		context.Background(),
		"UPDATE clientes SET saldo = $1 WHERE id = $2",
		novo_saldo,
		clienteId,
	)
	if err != nil {
		return nil, err
	}

	err = tx.Commit(context.Background())
	if err != nil {
		return nil, err
	}

	return &domain.TransacaoResponse{
		Limite: limite,
		Saldo:  novo_saldo,
	}, nil
}
