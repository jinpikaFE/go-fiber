package routers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	controller "github.com/jinpikaFE/go_fiber/controllers"
	"github.com/jinpikaFE/go_fiber/pkg/logging"
	"github.com/jinpikaFE/go_fiber/pkg/setting"

	"github.com/gofiber/fiber/v2/middleware/monitor"

	"github.com/gofiber/fiber/v2/middleware/cors"

	swagger "github.com/arsmn/fiber-swagger/v2"
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

	app.Static(
		"/docs",  // mount address
		"./docs", // path to the file folder
	)

	// swag init 命令生成
	app.Get("/docs/*", swagger.New(swagger.Config{ // custom
		URL:         fmt.Sprintf("http://localhost:%d/docs/swagger.json", setting.HTTPPort),
		DeepLinking: false,
	}))

	// // 跨域
	// app.Use(cors.New(cors.Config{
	// 	AllowOrigins: "http://localhost:8080",
	// }))

	// 跨域
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET,POST,PUT,DELETE,OPTIONS",
		AllowHeaders: "*",
	}))

	// 监控
	app.Get("/metrics", monitor.New())

	apiv1 := app.Group("/v1")

	{
		apiv1.Get("/region", controller.GetRegion)
	}

	{
		apiv1.Post("/mgb/monitor", controller.SetMgbData)
		apiv1.Get("/mgb/monitor", controller.GetMgbMonitor)
		apiv1.Get("/mgb/monitor/screen/:id", controller.GetMgbRecordScreen)
		apiv1.Get("/mgb/monitor/statist", controller.GetStatistRes)
	}

	{
		apiv1.Get("/test", controller.GetTests)
		apiv1.Post("/test", controller.AddTest)
		apiv1.Put("/test/:id", controller.EditTest)
		apiv1.Delete("/test/:id", controller.DelTest)
	}

	{
		apiv1.Post("/upload", controller.Upload)
		apiv1.Post("/uploadSource", controller.UploadSource)
	}

	return app
}
