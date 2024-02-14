package controllers

import (
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/kauefraga/esquilo-aniquilador/database"
	"github.com/kauefraga/esquilo-aniquilador/internal/domain"
)

func Transacoes(c *fiber.Ctx) error {
	clienteId, err := c.ParamsInt("id")
	if err != nil {
		return c.SendStatus(http.StatusNotFound)
	}

	// Verifica se o ID é de um cliente existente
	cliente, ok := database.Clientes[clienteId] // in-memory
	if !ok {
		return c.Status(http.StatusNotFound).JSON(&fiber.Map{
			"erro": "Cliente não existe.",
		})
	}

	transacao := domain.TransacaoPayload{}

	err = c.BodyParser(&transacao)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"erro": "O corpo da requisão é inválido.",
			"esperado": &fiber.Map{
				"valor":     100000,
				"tipo":      "c",
				"descricao": "descricao",
			},
			"observacoes": &fiber.Map{
				"valor":     "O valor deve estar em centavos e ser maior que 0.",
				"tipo":      "O tipo deve ser 'c' (crédito) ou 'd' (débito).",
				"descricao": "A descrição deve ter de 1 a 10 caractéres.",
			},
		})
	}

	// Verifica campos do payload
	if transacao.Valor <= 0 {
		return c.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"erro": "O valor da transação deve ser maior que zero.",
		})
	}

	if transacao.Tipo != "c" && transacao.Tipo != "d" {
		return c.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"erro": "A transação deve ser do tipo 'c' (crédito) ou 'd' (débito).",
		})
	}

	if len(transacao.Descricao) < 1 || len(transacao.Descricao) > 10 {
		return c.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"erro": "A descrição deve ter de 1 a 10 caractéres.",
		})
	}

	// Regra de negócio - Uma transação de débito NUNCA pode deixar o saldo do cliente menor que seu limite disponível
	// https://twitter.com/rkauefraga/status/1757524333629464861
	if transacao.Tipo == "d" && cliente.Saldo-transacao.Valor < -cliente.Limite {
		return c.Status(http.StatusUnprocessableEntity).JSON(&fiber.Map{
			"erro":       "A transação do tipo 'd' (débito) nunca pode deixar o saldo do cliente menor que seu limite disponível.",
			"observacao": "Transações do tipo 'c' (crédito) ainda serão processadas.",
		})
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

	return c.JSON(domain.TransacaoResponse{
		Limite: cliente.Limite,
		Saldo:  cliente.Saldo,
	})
}
