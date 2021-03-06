package routers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	jwtware "github.com/gofiber/jwt/v3"
	controller "github.com/jinpikaFE/go_fiber/controllers"
	"github.com/jinpikaFE/go_fiber/middleware/jwt"
	"github.com/jinpikaFE/go_fiber/pkg/logging"
	"github.com/jinpikaFE/go_fiber/pkg/untils"

	"github.com/gofiber/fiber/v2/middleware/monitor"
)

func nextLogger(c *fiber.Ctx) bool {
	// return true 会跳过本次中间件执行
	if c.Path() == "/v1/upload" {
		return true
	}
	return false
}

func stackTraceHandler(c *fiber.Ctx, e interface{}) {
	logging.Error(c, e)
}

func InitRouter() *fiber.App {
	app := fiber.New(fiber.Config{
		// ReadTimeout:     setting.ReadTimeout,
		// WriteTimeout:    setting.WriteTimeout,
		BodyLimit: 1000 * 1024 * 1024,
	})

	// panic 错误会被该中间件捕获
	app.Use(recover.New(recover.Config{
		EnableStackTrace:  true,
		StackTraceHandler: stackTraceHandler,
	}))

	app.Use(requestid.New())
	app.Use(logger.New(logger.Config{
		Next:   nextLogger,
		Format: "[INFO-${locals:requestid}]${time} pid: ${pid} status:${status} - ${method} path: ${path} queryParams: ${queryParams} body: ${body}\n resBody: ${resBody}\n error: ${error}\n",
		Output: logging.F,
	}))

	// 监控
	app.Get("/metrics", monitor.New())

	apiv1 := app.Group("/v1")

	{
		apiv1.Post("/login", controller.Login)
	}

	// JWT Middleware
	app.Use(jwtware.New(jwtware.Config{
		SigningKey:   []byte("secret"),
		ErrorHandler: untils.JwtError,
	}))

	apiv1.Use(jwt.Jwt)

	{
		apiv1.Get("/test", controller.GetTests)
		apiv1.Post("/test", controller.AddTest)
		apiv1.Put("/test/:id", controller.EditTest)
		apiv1.Delete("/test/:id", controller.DelTest)
	}

	{
		apiv1.Post("/upload", controller.Upload)
	}

	return app
}
