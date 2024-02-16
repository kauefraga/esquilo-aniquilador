package controllers

import (
	"errors"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/kauefraga/esquilo-aniquilador/internal/domain"
	"github.com/kauefraga/esquilo-aniquilador/internal/services"
)

func Transacoes(c *fiber.Ctx) error {
	// Pega param :id
	clienteId, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"erro": "O Cliente ID deve ser um número inteiro.",
		})
	}

	transacao := domain.TransacaoRequest{}

	// Parsa request body
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

	transacaoResponse, err := services.CreateTransacao(clienteId, &transacao)
	if err != nil {
		if errors.Is(err, services.ErroClienteNaoExiste) {
			return c.Status(http.StatusNotFound).JSON(&fiber.Map{
				"erro": "Cliente não existe.",
			})
		}

		if errors.Is(err, services.ErroValorDaTransacao) {
			return c.Status(http.StatusBadRequest).JSON(&fiber.Map{
				"erro": "O valor da transação deve ser maior que zero.",
			})
		}

		if errors.Is(err, services.ErroTipoDaTransacao) {
			return c.Status(http.StatusBadRequest).JSON(&fiber.Map{
				"erro": "A transação deve ser do tipo 'c' (crédito) ou 'd' (débito).",
			})
		}

		if errors.Is(err, services.ErroDescricao) {
			return c.Status(http.StatusBadRequest).JSON(&fiber.Map{
				"erro": "A descrição deve ter de 1 a 10 caractéres.",
			})
		}

		if errors.Is(err, services.ErroTransacaoDebito) {
			return c.Status(http.StatusUnprocessableEntity).JSON(&fiber.Map{
				"erro":       "A transação do tipo 'd' (débito) nunca pode deixar o saldo do cliente menor que seu limite disponível.",
				"observacao": "Transações do tipo 'c' (crédito) ainda serão processadas.",
			})
		}

		return c.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"erro": err.Error(),
		})
	}

	return c.JSON(transacaoResponse)
}
