package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
	_ "github.com/shaband/POS/docs"
	"go.uber.org/fx"
)

type Swagger struct {
	fx.In
	Route fiber.Router `name:"internal"`
}

func RegisterSwagger(c Swagger) {
	c.Route.Get("/swagger/*", swagger.HandlerDefault) // default
}
