package controllers

import (
	"errors"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/kauefraga/esquilo-aniquilador/internal/services"
)

func Extrato(c *fiber.Ctx) error {
	clienteId, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"erro": "O Cliente ID deve ser um número inteiro.",
		})
	}

	extrato, err := services.GetExtrato(clienteId)
	if err != nil {
		if errors.Is(err, services.ErroClienteNaoExiste) {
			return c.Status(http.StatusNotFound).JSON(&fiber.Map{
				"erro": "Cliente não existe.",
			})
		}

		return c.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"erro": err.Error(),
		})
	}

	return c.JSON(extrato)
}
