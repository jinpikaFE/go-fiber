package controller

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/jinpikaFE/go_fiber/models"
	"github.com/jinpikaFE/go_fiber/pkg/app"
	"github.com/jinpikaFE/go_fiber/pkg/e"
	"github.com/jinpikaFE/go_fiber/pkg/gredis"
	"github.com/jinpikaFE/go_fiber/pkg/logging"
	"github.com/jinpikaFE/go_fiber/pkg/untils"
	"github.com/jinpikaFE/go_fiber/pkg/valodates"
)

// 获取User列表
func GetUsers(c *fiber.Ctx) error {
	// maps := make(map[string]interface{})
	// // 获取get query参数 或者使用queryparser
	// id := c.Query("id")
	// maps["id"] = id
	appF := app.Fiber{C: c}
	maps := &models.User{}
	page := &e.PageStruct{}
	err2 := c.QueryParser(page)

	if err := c.QueryParser(maps); err != nil && err2 != nil {
		logging.Error(err)
		return appF.Response(fiber.StatusInternalServerError, fiber.StatusInternalServerError, "参数解析错误", err)
	}

	// 入参验证
	if errors := valodates.ValidateStruct(*page); errors != nil {
		return appF.Response(fiber.StatusBadRequest, fiber.StatusBadRequest, "检验参数错误", errors)
	}

	res, errs := models.GetUsers((page.Page-1)*page.PageSize, page.PageSize, maps)
	if errs != nil {
		return appF.Response(fiber.StatusInternalServerError, fiber.StatusInternalServerError, "查询失败", errs)
	}

	total, errTotal := models.GetUserTotal(maps)
	if errTotal != nil {
		return appF.Response(fiber.StatusInternalServerError, fiber.StatusInternalServerError, "查询失败", errTotal)
	}

	data := make(map[string]interface{})
	data["list"] = res
	data["total"] = total
	data["pageSize"] = page.PageSize
	data["page"] = page.Page

	return appF.Response(fiber.StatusOK, fiber.StatusOK, "SUCCESS", data)
}

// 获取User
func GetUser(c *fiber.Ctx) error {
	// maps := make(map[string]interface{})
	// // 获取get query参数 或者使用queryparser
	// id := c.Query("id")
	// maps["id"] = id
	appF := app.Fiber{C: c}
	maps := &models.User{}

	if err := c.QueryParser(maps); err != nil {
		logging.Error(err)
		return appF.Response(fiber.StatusInternalServerError, fiber.StatusInternalServerError, "参数解析错误", err)
	}

	res, errs := models.GetUser(maps)

	if errs != nil {
		return appF.Response(fiber.StatusInternalServerError, fiber.StatusInternalServerError, "查询失败", errs)
	}

	if !(res.ID > 0) {
		return appF.Response(fiber.StatusBadRequest, fiber.StatusBadRequest, "用户不存在", nil)
	}

	return appF.Response(fiber.StatusOK, fiber.StatusOK, "SUCCESS", res)
}

// 添加user
func AddUser(c *fiber.Ctx) error {
	appF := app.Fiber{C: c}
	user := &models.User{}

	if err := c.BodyParser(user); err != nil {
		return appF.Response(fiber.StatusInternalServerError, fiber.StatusInternalServerError, "参数解析错误", err)
	}

	// 入参验证
	errors := valodates.ValidateStruct(*user)

	if errors != nil {
		return appF.Response(fiber.StatusBadRequest, fiber.StatusBadRequest, "检验参数错误", errors)
	}

	user.Password = untils.GetSha256(user.Password)

	err := models.AddUser(*user)

	if err != nil {
		return appF.Response(fiber.StatusInternalServerError, fiber.StatusInternalServerError, "添加失败", err)
	}

	return appF.Response(fiber.StatusOK, fiber.StatusOK, "SUCCESS", nil)
}

// 编辑用户
func EditUser(c *fiber.Ctx) error {
	appF := app.Fiber{C: c}
	id, err := strconv.Atoi(c.Params("id"))
	user := &models.User{}

	if err := c.BodyParser(user); err != nil {
		logging.Error(err)
		return appF.Response(fiber.StatusInternalServerError, fiber.StatusInternalServerError, "参数解析错误", err)
	}

	// 如果手机号存在走手机号修改逻辑，验证验证码
	if user.Mobile != nil {
		loginMobile := &models.LoginMobile{}
		if err := c.BodyParser(loginMobile); err != nil {
			logging.Error(err)
			return appF.Response(fiber.StatusInternalServerError, fiber.StatusInternalServerError, "参数解析错误", err)
		}

		// 入参校验
		if err := valodates.ValidateStruct(*loginMobile); err != nil {
			return appF.Response(fiber.StatusBadRequest, fiber.StatusBadRequest, "检验参数错误", err)
		}
		if models.ExistUserByID(id) {
			reply, replyErr := gredis.Get(loginMobile.Mobile)
			if replyErr != nil {
				logging.Error(replyErr)
				return appF.Response(fiber.StatusInternalServerError, fiber.StatusInternalServerError, "redis错误", replyErr)
			}

			if loginMobile.Captcha == string(reply) {
				if err := models.EditUser(id, *user); err != nil {
					return appF.Response(fiber.StatusInternalServerError, fiber.StatusInternalServerError, "编辑失败", err)
				}
				return appF.Response(fiber.StatusOK, fiber.StatusOK, "SUCCESS", nil)
			}

			return appF.Response(fiber.StatusBadRequest, fiber.StatusBadRequest, "验证码错误", nil)
		} else {
			return appF.Response(fiber.StatusBadRequest, fiber.StatusBadRequest, "id不存在", nil)
		}

	}

	if err != nil {
		logging.Error(err)
		return appF.Response(fiber.StatusInternalServerError, fiber.StatusInternalServerError, "Params： id 参数解析错误", err)
	}
	if models.ExistUserByID(id) {
		if err := models.EditUser(id, *user); err != nil {
			return appF.Response(fiber.StatusInternalServerError, fiber.StatusInternalServerError, "编辑失败", err)
		}
		return appF.Response(fiber.StatusOK, fiber.StatusOK, "SUCCESS", nil)
	}

	return appF.Response(fiber.StatusBadRequest, fiber.StatusBadRequest, "id不存在", nil)
}

// 删除user
func DelUser(c *fiber.Ctx) error {
	appF := app.Fiber{C: c}
	id, err := strconv.Atoi(c.Params("id"))

	if err != nil {
		logging.Error(err)
		return appF.Response(fiber.StatusInternalServerError, fiber.StatusInternalServerError, "Params： id 参数解析错误", err)
	}

	if models.ExistUserByID(id) {
		if err := models.DeleteUser(id); err != nil {
			return appF.Response(fiber.StatusInternalServerError, fiber.StatusInternalServerError, "删除失败", err)
		}
		return appF.Response(fiber.StatusOK, fiber.StatusOK, "SUCCESS", nil)
	}

	return appF.Response(fiber.StatusBadRequest, fiber.StatusBadRequest, "id不存在", nil)

}
