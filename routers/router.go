package routers

import (
	"github.com/gofiber/fiber/v2"
	v1 "github.com/jinpikaFE/go_fiber/routers/v1"
)

func InitRouter() *fiber.App {
	app := fiber.New()

	apiv1 := app.Group("/v1")

	{
		apiv1.Get("/test", v1.GetTests)
	}

	return app
}
