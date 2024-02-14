package controllers

import (
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/kauefraga/esquilo-aniquilador/database"
	"github.com/kauefraga/esquilo-aniquilador/internal/domain"
)

func Extrato(c *fiber.Ctx) error {
	clienteId, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"erro": "O Cliente ID deve ser um número inteiro.",
		})
	}

	// Verifica se o ID é de um cliente existente
	cliente, ok := database.Clientes[clienteId] // in-memory
	if !ok {
		return c.Status(http.StatusNotFound).JSON(&fiber.Map{
			"erro": "Cliente não existe.",
		})
	}

	// Pegando últimas transações
	var ultimasTransacoes []domain.UltimaTransacao
	for _, t := range database.Transacoes {
		if t.ClienteId == cliente.ID {
			if len(ultimasTransacoes) == 10 {
				break
			} // Até 10 registros

			ultimasTransacoes = append(
				ultimasTransacoes,
				domain.UltimaTransacao{
					Valor:       t.Valor,
					Tipo:        t.Tipo,
					Descricao:   t.Descricao,
					RealizadaEm: t.RealizadaEm,
				},
			)
		}
	} // in-memory

	// Sem ordenar
	return c.JSON(&domain.Extrato{
		Saldo: domain.Saldo{
			Total:       cliente.Saldo,
			DataExtrato: time.Now(),
			Limite:      cliente.Limite,
		},
		UltimasTransacoes: ultimasTransacoes,
	})
}
