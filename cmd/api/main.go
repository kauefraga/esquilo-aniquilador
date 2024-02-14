package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kauefraga/esquilo-aniquilador/database"
	"github.com/kauefraga/esquilo-aniquilador/internal/controllers"
	"github.com/kauefraga/esquilo-aniquilador/internal/domain"
)

func main() {
	app := fiber.New()

	database.Clientes = make(map[int]domain.Cliente)
	database.Transacoes = make(map[int]domain.Transacao)
	database.InsereClientesIniciais()

	app.Get("/", controllers.Hello)
	app.Get("/clientes/:id/extrato", controllers.Extrato)
	app.Post("/clientes/:id/transacoes", controllers.Transacoes)

	app.Listen(":3000")
}
