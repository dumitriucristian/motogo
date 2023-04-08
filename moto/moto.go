package moto

import (
	"github.com/gofiber/fiber"
)

func GetMotors(c *fiber.Ctx) {
	c.Send("All motorcycle")
}

func GetMoto(c *fiber.Ctx) {
	c.Send("A motorcycle")
}

func NewMoto(c *fiber.Ctx) {
	c.Send("A new motorcycle")
}

func DeleteMoto(c *fiber.Ctx) {
	c.Send("Delete motorcycles")
}
