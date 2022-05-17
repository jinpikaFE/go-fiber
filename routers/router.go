package routers

import (
	"github.com/gofiber/fiber/v2"
	controller "github.com/jinpikaFE/go_fiber/controllers"
)

func InitRouter() *fiber.App {
	app := fiber.New()

	apiv1 := app.Group("/v1")

	{
		apiv1.Get("/test", controller.GetTests)
		apiv1.Post("/test", controller.AddTest)
		apiv1.Put("/test/:id", controller.EditTest)
		apiv1.Delete("/test/:id", controller.DelTest)
	}

	return app
}
