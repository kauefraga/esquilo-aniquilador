package main

import (
	"github.com/gofiber/contrib/swagger"
	"github.com/gofiber/fiber/v2"
	"github.com/kauefraga/esquilo-aniquilador/internal/controllers"
)

func main() {
	app := fiber.New()

	app.Use(swagger.New())

	app.Get("/", controllers.Hello)
	app.Get("/clientes/:id/extrato", controllers.Extrato)
	app.Post("/clientes/:id/transacoes", controllers.Transacoes)

	app.Listen(":3000")
}
