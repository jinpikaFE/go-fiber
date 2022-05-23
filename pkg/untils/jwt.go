package untils

import "github.com/gofiber/fiber/v2"

func JwtError(c *fiber.Ctx, err error) error {
	if err.Error() == "Missing or malformed JWT" {
		return c.Status(fiber.StatusBadRequest).
			JSON(fiber.Map{"code": fiber.StatusBadRequest, "message": "Missing or malformed JWT", "data": nil})
	}
	return c.Status(fiber.StatusUnauthorized).
		JSON(fiber.Map{"code": fiber.StatusUnauthorized, "message": "Invalid or expired JWT", "data": nil})
}
