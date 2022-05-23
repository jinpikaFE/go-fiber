package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jinpikaFE/go_fiber/models"
	"github.com/jinpikaFE/go_fiber/pkg/logging"
	"github.com/jinpikaFE/go_fiber/pkg/valodates"
)

// 添加test
func Login(c *fiber.Ctx) error {
	login := &models.Login{}
	code := 200
	message := "SUCCESS"

	if err := c.BodyParser(login); err != nil {
		code = 500
		message = "参数解析错误"
		logging.Error(err)
	}

	// 入参验证
	errors := valodates.ValidateStruct(*login)

	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"code":    code,
			"message": message,
			"data":    errors,
		})
	}

	res := models.GetToken(login)

	if res == "" {
		code = fiber.StatusUnauthorized
		message = "账户或者密码错误"
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"code":    code,
			"message": message,
			"data":    res,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"code":    code,
		"message": message,
		"data":    res,
	})
}
