package main

import (
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/kauefraga/esquilo-aniquilador/internal/controllers"
)

func main() {
	app := fiber.New()

	app.Get("/clientes/:id/extrato", controllers.Extrato)
	app.Post("/clientes/:id/transacoes", controllers.Transacoes)

	port := fmt.Sprintf(":%s", os.Getenv("PORT"))

	app.Listen(port)
}
