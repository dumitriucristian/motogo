package main

import (
	"github.com/gofiber/fiber"
	restmoto/packages/moto
)

func helloWorld(c *fiber.Ctx) {
	c.Send("Hello, World")
}

func setupRoutes(app *fiber.App) {
	app.Get("/", helloWorld)

	app.Get("api/v1/motors", moto.GetMotors)
	app.Get("api/v1/moto/:id", moto.GetMoto)
	app.Post("api/v1/moto", moto.AddMoto)
	app.Delete("api/v1/motors", moto.DeleteMoto)
}

func main() {
	app := fiber.New()

	setupRoutes(app)
	app.Listen(3000)
}
