package controller

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/url"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/jinpikaFE/go_fiber/models"
	"github.com/jinpikaFE/go_fiber/pkg/app"
	"github.com/jinpikaFE/go_fiber/pkg/gredis"
	"github.com/jinpikaFE/go_fiber/pkg/logging"
	"github.com/jinpikaFE/go_fiber/pkg/tencent"
	"github.com/jinpikaFE/go_fiber/pkg/untils"
	"github.com/jinpikaFE/go_fiber/pkg/valodates"
	"github.com/vicanso/go-axios"
)

// ResponseHTTP represents response body of this API
type ResponseHTTP struct {
	Code    int         `json:"code"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

// 登录
// @Summary 登录接口
// @Description 登录接口
// @Tags 登录
// @Accept json
// @Produce json
// @Param login body models.Login true "Login"
// @Success 200 {object} ResponseHTTP{}
// @Failure 503 {object} ResponseHTTP{}
// @Router /v1/login [post]
func Login(c *fiber.Ctx) error {
	appF := app.Fiber{C: c}
	types := &models.Type{}

	if err := c.BodyParser(types); err != nil {
		return appF.Response(fiber.StatusInternalServerError, fiber.StatusInternalServerError, "参数解析错误", nil)
	}

	// 账号登录
	if types.LoginType == "1" {
		loginAccount := &models.LoginAccount{}

		if err := c.BodyParser(loginAccount); err != nil {
			return appF.Response(fiber.StatusInternalServerError, fiber.StatusInternalServerError, "参数解析错误", nil)
		}
		// 入参验证
		errors := valodates.ValidateStruct(*loginAccount)

		if errors != nil {
			return appF.Response(fiber.StatusBadRequest, fiber.StatusBadRequest, "检验参数错误", errors)
		}

		userSt := &models.User{}
		userSt.Username = &loginAccount.Username

		res, errs := models.GetUser(userSt)

		if errs != nil {
			return appF.Response(fiber.StatusInternalServerError, fiber.StatusInternalServerError, "查询失败", errs)
		}

		if !(res.ID > 0) {
			return appF.Response(fiber.StatusBadRequest, fiber.StatusBadRequest, "用户不存在", nil)
		}

		login := &models.Login{}
		login.Username = loginAccount.Username
		login.Password = loginAccount.Password

		token := models.GetToken(login, res)

		redisErr := gredis.Set("token", token, 300, true)

		if redisErr != nil {
			logging.Error(redisErr)
			return appF.Response(fiber.StatusInternalServerError, fiber.StatusInternalServerError, "redis错误", redisErr)
		}

		if token == "" {
			return appF.Response(fiber.StatusUnauthorized, fiber.StatusUnauthorized, "账户或者密码错误", nil)
		}

		loginres := map[string]interface{}{"token": token, "username": loginAccount.Username}

		return appF.Response(fiber.StatusOK, fiber.StatusOK, "SUCCESS", loginres)
	}

	// 微信登录
	if types.LoginType == "2" {
		loginWx := &models.LoginWx{}

		if err := c.BodyParser(loginWx); err != nil {
			return appF.Response(fiber.StatusInternalServerError, fiber.StatusInternalServerError, "参数解析错误", nil)
		}

		// 入参验证
		errors := valodates.ValidateStruct(*loginWx)

		if errors != nil {
			return appF.Response(fiber.StatusBadRequest, fiber.StatusBadRequest, "检验参数错误", errors)
		}
		// 使用axios进行请求
		queryParams := url.Values{}
		queryParams.Add("appid", loginWx.Appid)
		queryParams.Add("secret", loginWx.Appsecret)
		queryParams.Add("js_code", loginWx.Code)
		queryParams.Add("grant_type", "authorization_code")
		axiosConfig := &axios.InstanceConfig{}
		axiosConfig.BaseURL = "https://api.weixin.qq.com"
		resp, err := untils.Request(axiosConfig).Get("/sns/jscode2session", queryParams)
		if err != nil {
			return appF.Response(fiber.StatusInternalServerError, fiber.StatusInternalServerError, "获取openid失败", err)
		}
		result := make(map[string]interface{})
		if errJson := json.Unmarshal(resp.Data, &result); errJson != nil {
			return appF.Response(fiber.StatusInternalServerError, fiber.StatusInternalServerError, "解析code2Session数据失败", errJson)
		}
		logging.Error(result, result["openid"])
		// body, err := ioutil.ReadAll(response.Body)
		// if err == nil {
		// 	logging.Error(string(body))
		// }

		userSt := &models.User{}
		openid := result["openid"].(string)
		userSt.Openid = &openid
		resUser, errs := models.GetUser(userSt)
		if errs != nil {
			return appF.Response(fiber.StatusInternalServerError, fiber.StatusInternalServerError, "查询失败", errs)
		}

		userWx := &models.User{}
		userWx.Openid = &openid

		if !(resUser.ID > 0) {
			// 不存在就创建

			nickName := "微信用户"
			userWx.NickName = &nickName
			if err := models.AddUser(*userWx); err != nil {
				return appF.Response(fiber.StatusInternalServerError, fiber.StatusInternalServerError, "添加失败", err)
			}
		}
		// 登录

		token := models.GetToken(&models.Login{}, userWx)

		redisErr := gredis.Set("token", token, 300, true)

		if redisErr != nil {
			logging.Error(redisErr)
			return appF.Response(fiber.StatusInternalServerError, fiber.StatusInternalServerError, "redis错误", redisErr)
		}

		loginres := map[string]interface{}{"token": token, "openid": openid}

		return appF.Response(fiber.StatusOK, fiber.StatusOK, "SUCCESS", loginres)
	}

	// 手机号登录
	if types.LoginType == "3" {
		loginMobile := &models.LoginMobile{}

		if err := c.BodyParser(loginMobile); err != nil {
			return appF.Response(fiber.StatusInternalServerError, fiber.StatusInternalServerError, "参数解析错误", nil)
		}

		// 入参验证
		errors := valodates.ValidateStruct(*loginMobile)

		if errors != nil {
			return appF.Response(fiber.StatusBadRequest, fiber.StatusBadRequest, "检验参数错误", errors)
		}

		userSt := &models.User{}
		userSt.Mobile = &loginMobile.Mobile
		resUser, errs := models.GetUser(userSt)
		if errs != nil {
			return appF.Response(fiber.StatusInternalServerError, fiber.StatusInternalServerError, "查询失败", errs)
		}

		userMobile := &models.User{}
		userMobile.Mobile = &loginMobile.Mobile

		if !gredis.Exists(loginMobile.Mobile) {
			return appF.Response(fiber.StatusBadRequest, fiber.StatusBadRequest, "验证码过期请重新获取", nil)
		}
		reply, replyErr := gredis.Get(loginMobile.Mobile)
		if replyErr != nil {
			logging.Error(replyErr)
			return appF.Response(fiber.StatusInternalServerError, fiber.StatusInternalServerError, "redis错误", replyErr)
		}

		if loginMobile.Captcha == string(reply) {
			if !(resUser.ID > 0) {
				// 不存在就创建
				userMobile.NickName = &loginMobile.Mobile
				if err := models.AddUser(*userMobile); err != nil {
					return appF.Response(fiber.StatusInternalServerError, fiber.StatusInternalServerError, "添加失败", err)
				}
			}
			// 登录

			token := models.GetToken(&models.Login{}, userMobile)
			redisErr := gredis.Set("token", token, 300, true)

			if redisErr != nil {
				logging.Error(redisErr)
				return appF.Response(fiber.StatusInternalServerError, fiber.StatusInternalServerError, "redis错误", redisErr)
			}

			loginres := map[string]interface{}{"token": token, "mobile": &loginMobile.Mobile}

			return appF.Response(fiber.StatusOK, fiber.StatusOK, "SUCCESS", loginres)
		}

		return appF.Response(fiber.StatusBadRequest, fiber.StatusBadRequest, "验证码错误", nil)
	}

	return appF.Response(fiber.StatusBadRequest, fiber.StatusBadRequest, "未知登录类型", nil)
}

// 获取验证码
// @Summary 获取验证码
// @Description 获取验证码
// @Tags 获取验证码
// @Accept json
// @Produce json
// @Param mobile body models.LoginMobile.Mobile true "mobile"
// @Success 200 {object} ResponseHTTP{}
// @Failure 503 {object} ResponseHTTP{}
// @Router /v1/captcha [post]
func GetCaptcha(c *fiber.Ctx) error {
	appF := app.Fiber{C: c}
	// 短信验证码发送
	loginMobile := &models.LoginMobile{}

	captcha := fmt.Sprintf("%v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(1000000))

	if err := c.BodyParser(loginMobile); err != nil {
		return appF.Response(fiber.StatusInternalServerError, fiber.StatusInternalServerError, "参数解析错误", nil)
	}

	if !untils.VerifyMobileFormat(loginMobile.Mobile) {
		return appF.Response(fiber.StatusBadRequest, fiber.StatusBadRequest, "请输入正确的手机号", nil)
	}

	result, smsErr := tencent.SendSms(captcha, loginMobile.Mobile)

	if smsErr != nil {
		return appF.Response(fiber.StatusInternalServerError, fiber.StatusInternalServerError, "短信发送错误", smsErr)
	}

	if result.SendStatusSet[0].Code == "Ok" {
		redisErr := gredis.Set(loginMobile.Mobile, captcha, 300, true)

		if redisErr != nil {
			logging.Error(redisErr)
			return appF.Response(fiber.StatusInternalServerError, fiber.StatusInternalServerError, "redis错误", redisErr)
		}
		return appF.Response(fiber.StatusOK, fiber.StatusOK, "短信发送成功", result)
	}

	return appF.Response(fiber.StatusInternalServerError, fiber.StatusInternalServerError, "短信发送失败", result)
}
