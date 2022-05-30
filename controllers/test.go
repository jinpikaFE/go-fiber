package controller

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/jinpikaFE/go_fiber/models"
	"github.com/jinpikaFE/go_fiber/pkg/app"
	"github.com/jinpikaFE/go_fiber/pkg/logging"
	"github.com/jinpikaFE/go_fiber/pkg/valodates"
)

// 获取Test列表
func GetTests(c *fiber.Ctx) error {
	// maps := make(map[string]interface{})
	// // 获取get query参数 或者使用queryparser
	// id := c.Query("id")
	// maps["id"] = id
	appF := app.Fiber{C: c}
	maps := &models.Test{}
	if err := c.QueryParser(maps); err != nil {
		logging.Error(err)
		return appF.Response(fiber.StatusInternalServerError, fiber.StatusInternalServerError, "参数解析错误", err)
	}
	res := models.GetTests(0, 10, maps)

	return appF.Response(fiber.StatusOK, fiber.StatusOK, "SUCCESS", res)
}

// 添加test
func AddTest(c *fiber.Ctx) error {
	appF := app.Fiber{C: c}
	test := &models.Test{}

	if err := c.BodyParser(test); err != nil {
		logging.Error(err)
		return appF.Response(fiber.StatusInternalServerError, fiber.StatusInternalServerError, "参数解析错误", err)
	}

	// 入参验证
	errors := valodates.ValidateStruct(*test)

	if errors != nil {
		return appF.Response(fiber.StatusBadRequest, fiber.StatusBadRequest, "检验参数错误", errors)
	}

	res := models.AddTest(test)

	return appF.Response(fiber.StatusOK, fiber.StatusOK, "SUCCESS", res)
}

// 编辑test
func EditTest(c *fiber.Ctx) error {
	appF := app.Fiber{C: c}
	id, err := strconv.Atoi(c.Params("id"))
	test := &models.Test{}
	res := false

	if err := c.BodyParser(test); err != nil {
		logging.Error(err)
		return appF.Response(fiber.StatusInternalServerError, fiber.StatusInternalServerError, "参数解析错误", err)
	}

	if err != nil {
		logging.Error(err)
		return appF.Response(fiber.StatusInternalServerError, fiber.StatusInternalServerError, "Params： id 参数解析错误", err)
	}
	if models.ExistTestByID(id) {
		res = models.EditTest(id, test)
	} else {
		return appF.Response(fiber.StatusBadRequest, fiber.StatusBadRequest, "id不存在", nil)
	}

	return appF.Response(fiber.StatusOK, fiber.StatusOK, "SUCCESS", res)
}

// 删除test
func DelTest(c *fiber.Ctx) error {
	appF := app.Fiber{C: c}
	id, err := strconv.Atoi(c.Params("id"))
	res := false

	if err != nil {
		logging.Error(err)
		return appF.Response(fiber.StatusInternalServerError, fiber.StatusInternalServerError, "Params： id 参数解析错误", err)
	}

	if models.ExistTestByID(id) {
		res = models.DeleteTest(id)
	} else {
		return appF.Response(fiber.StatusBadRequest, fiber.StatusBadRequest, "id不存在", nil)
	}

	return appF.Response(fiber.StatusOK, fiber.StatusOK, "SUCCESS", res)
}
