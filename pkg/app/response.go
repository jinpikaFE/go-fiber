package app

import "github.com/gofiber/fiber/v2"

type Fiber struct {
	C *fiber.Ctx
}

func (g *Fiber) Response(httpCode, errCode int, message string, data interface{}) error {

	return g.C.Status(fiber.StatusOK).JSON(fiber.Map{
		"code":    errCode,
		"message": message,
		"data":    data,
	})
}
