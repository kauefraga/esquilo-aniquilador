package services

import (
	"errors"
	"time"

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
	// Verifica se o ID é de um cliente existente
	cliente, ok := database.Clientes[clienteId] // in-memory
	if !ok {
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

	// Verifica regra de negócio
	// Regra de negócio - Uma transação de débito NUNCA pode deixar o saldo do cliente menor que seu limite disponível
	// https://twitter.com/rkauefraga/status/1757524333629464861
	if transacao.Tipo == "d" && cliente.Saldo-transacao.Valor < -cliente.Limite {
		return nil, ErroTransacaoDebito
	}

	// Para pegar o último elemento: `len(database.Transacoes)-1`.
	// O seguinte seria `len(database.Transacoes)-1+1`
	database.Transacoes[len(database.Transacoes)] = domain.Transacao{
		ID:          len(database.Transacoes),
		Valor:       transacao.Valor,
		Tipo:        transacao.Tipo,
		Descricao:   transacao.Descricao,
		ClienteId:   cliente.ID,
		RealizadaEm: time.Now(),
	} // in-memory

	cliente.Saldo -= transacao.Valor
	database.Clientes[cliente.ID] = cliente

	return &domain.TransacaoResponse{
		Limite: cliente.Limite,
		Saldo:  cliente.Saldo,
	}, nil
}
