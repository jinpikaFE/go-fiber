package models

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/jinpikaFE/go_fiber/pkg/logging"
	"github.com/jinpikaFE/go_fiber/pkg/untils"
)

type LoginAccount struct {
	Username string `validate:"required" query:"username" json:"username" xml:"username" form:"username"`
	Password string `validate:"required" query:"password" json:"password" xml:"password" form:"password"`
}

type LoginMobile struct {
	Mobile  string `validate:"required" query:"mobile" json:"mobile" xml:"mobile" form:"mobile"`
	Captcha string `validate:"required" query:"captcha" json:"captcha" xml:"captcha" form:"captcha"`
}

type LoginWx struct {
	Appid     string `validate:"required" query:"appid" json:"appid" xml:"appid" form:"appid"`
	Appsecret string `validate:"required" query:"appsecret" json:"appsecret" xml:"appsecret" form:"appsecret"`
	Code      string `validate:"required" query:"code" json:"code" xml:"code" form:"code"`
}

type Type struct {
	LoginType string `validate:"required,oneof=1 2 3" query:"loginType" json:"loginType" xml:"loginType" form:"loginType"`
}

type Login struct {
	LoginAccount

	LoginWx

	Type
}

func GetToken(login *Login, user *User) string {
	claims := jwt.MapClaims{
		"username": login.Username,
		"admin":    true,
		"exp":      time.Now().Add(time.Hour * 72).Unix(),
	}
	// logging.Error(user.Openid, user.Openid == nil)
	if user.Openid == nil {
		if user.Mobile == nil {
			if login.Username != *user.Username || untils.GetSha256(login.Password) != user.Password {
				return ""
			}
		} else {
			claims = jwt.MapClaims{
				"openid": user.Mobile,
				"admin":  true,
				"exp":    time.Now().Add(time.Hour * 72).Unix(),
			}
		}
	} else {
		claims = jwt.MapClaims{
			"openid": user.Openid,
			"admin":  true,
			"exp":    time.Now().Add(time.Hour * 72).Unix(),
		}
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		logging.Error(err)
		return ""
	}

	return t
}
