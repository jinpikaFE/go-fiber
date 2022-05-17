package controller

import (
	"log"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/jinpikaFE/go_fiber/models"
)

// var (
// 	me
// )

// 获取Test列表
func GetTests(c *fiber.Ctx) error {
	// maps := make(map[string]interface{})
	// // 获取get query参数 或者使用queryparser
	// id := c.Query("id")
	// maps["id"] = id

	maps := &models.Test{}
	c.QueryParser(maps)

	res := models.GetTests(0, 10, maps)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"code":    200,
		"message": "SUCCESS",
		"data":    res,
		"query":   maps,
	})
}

// 添加test
func AddTest(c *fiber.Ctx) error {
	test := &models.Test{}
	c.BodyParser(test)

	res := models.AddTest(test)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"code":    200,
		"message": "SUCCESS",
		"data":    res,
	})
}

// 编辑test
func EditTest(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	test := &models.Test{}
	c.BodyParser(test)
	code := 200
	res := false
	message := "SUCCESS"

	if err != nil {
		log.Fatal(err)
	}
	if models.ExistTestByID(id) {
		res = models.EditTest(id, test)
	} else {
		code = 500
	}

	if res {
		message = "SUCCESS"
	} else {
		message = "ERROR"
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
		log.Fatal(err)
	}

	if models.ExistTestByID(id) {
		res = models.DeleteTest(id)
	} else {
		code = 500
	}

	if res {
		message = "SUCCESS"
	} else {
		message = "ERROR"
	}


	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"code":    code,
		"message": message,
		"data":    res,
	})
}
