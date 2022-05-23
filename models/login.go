package models

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/jinpikaFE/go_fiber/pkg/logging"
)

type Login struct {
	Username string `validate:"required" query:"username" json:"username" xml:"username" form:"username"`
	Password string `validate:"required" query:"password" json:"password" xml:"password" form:"password"`
}

func GetToken(login *Login) string {
	if login.Username != "admin" || login.Password != "admin" {
		return ""
	}

	claims := jwt.MapClaims{
		"username": login.Username,
		"admin":    true,
		"exp":      time.Now().Add(time.Hour * 72).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		logging.Error(err)
		return ""
	}

	return t
}
