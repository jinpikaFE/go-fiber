package controller

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/jinpikaFE/go_fiber/models"
	"github.com/jinpikaFE/go_fiber/pkg/logging"
	"github.com/jinpikaFE/go_fiber/pkg/valodates"
)

// 获取Test列表
func GetTests(c *fiber.Ctx) error {
	// maps := make(map[string]interface{})
	// // 获取get query参数 或者使用queryparser
	// id := c.Query("id")
	// maps["id"] = id
	code := 200
	message := "SUCCESS"
	maps := &models.Test{}
	if err := c.QueryParser(maps); err != nil {
		code = 500
		message = "参数解析错误"
		logging.Error(err)
	}
	res := models.GetTests(0, 10, maps)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"code":    code,
		"message": message,
		"data":    res,
	})
}

// 添加test
func AddTest(c *fiber.Ctx) error {
	test := &models.Test{}
	code := 200
	message := "SUCCESS"

	if err := c.BodyParser(test); err != nil {
		code = 500
		message = "参数解析错误"
		logging.Error(err)
	}

	// 入参验证
	errors := valodates.ValidateStruct(*test)

	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"code":    code,
			"message": message,
			"data":    errors,
		})
	}

	res := models.AddTest(test)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"code":    code,
		"message": message,
		"data":    res,
	})
}

// 编辑test
func EditTest(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	test := &models.Test{}
	code := 200
	res := false
	message := "SUCCESS"

	if err := c.BodyParser(test); err != nil {
		code = 500
		message = "参数解析错误"
		logging.Error(err)
	}

	if err != nil {
		code = 500
		message = "Params： id 参数解析错误"
		logging.Error(err)
	}
	if models.ExistTestByID(id) {
		res = models.EditTest(id, test)
	} else {
		code = 500
		message = "id不存在"
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"code":    code,
		"message": message,
		"data":    res,
	})
}

// 删除test
func DelTest(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	code := 200
	res := false
	message := "SUCCESS"

	if err != nil {
		code = 500
		message = "Params： id 参数解析错误"
		logging.Error(err)
	}

	if models.ExistTestByID(id) {
		res = models.DeleteTest(id)
	} else {
		code = 500
		message = "id不存在"
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"code":    code,
		"message": message,
		"data":    res,
	})
}
