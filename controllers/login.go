package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jinpikaFE/go_fiber/models"
	"github.com/jinpikaFE/go_fiber/pkg/app"
	"github.com/jinpikaFE/go_fiber/pkg/gredis"
	"github.com/jinpikaFE/go_fiber/pkg/logging"
	"github.com/jinpikaFE/go_fiber/pkg/valodates"
)

// 添加test
func Login(c *fiber.Ctx) error {
	appF := app.Fiber{C: c}
	login := &models.Login{}

	if err := c.BodyParser(login); err != nil {
		logging.Error(err)
		return appF.Response(fiber.StatusInternalServerError, fiber.StatusInternalServerError, "参数解析错误", nil)
	}

	// 入参验证
	errors := valodates.ValidateStruct(*login)

	if errors != nil {
		logging.Error(errors)
		return appF.Response(fiber.StatusBadRequest, fiber.StatusBadRequest, "检验参数错误", errors)
	}

	res := models.GetToken(login)

	redisErr := gredis.Set("token", res, 300)

	if redisErr != nil {
		logging.Error(redisErr)
		return appF.Response(fiber.StatusInternalServerError, fiber.StatusInternalServerError, "redis错误", redisErr)
	}

	if res == "" {
		return appF.Response(fiber.StatusUnauthorized, fiber.StatusUnauthorized, "账户或者密码错误", nil)
	}

	return appF.Response(fiber.StatusOK, fiber.StatusOK, "SUCCESS", res)
}
