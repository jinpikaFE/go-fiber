package jwt

import (
	// "strings"
	// "time"

	"github.com/gofiber/fiber/v2"
	"github.com/jinpikaFE/go_fiber/pkg/logging"

	"github.com/golang-jwt/jwt/v4"
)

func Jwt(ctx *fiber.Ctx) error {
	// gofiber不需要去请求头中寻找Authorization: Bearer
	// 直接通过
	// user := c.Locals("user").(*jwt.Token)
	// claims := user.Claims.(jwt.MapClaims)
	user := ctx.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)

	logging.Info(claims)
	return ctx.Next()
}
