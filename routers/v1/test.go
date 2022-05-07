package v1

import "github.com/gofiber/fiber/v2"

// 获取Test列表
func GetTests(c *fiber.Ctx) error {

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "test",
	})
}
