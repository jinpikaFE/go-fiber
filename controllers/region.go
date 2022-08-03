package controller

import (
	"encoding/json"
	"net/url"

	"github.com/gofiber/fiber/v2"
	"github.com/jinpikaFE/go_fiber/pkg/app"
	"github.com/jinpikaFE/go_fiber/pkg/untils"
	"github.com/vicanso/go-axios"
)

// 获取省市区
// @Summary 获取省市区
// @Description 获取省市区
// @Tags 获取省市区
// @Accept json
// @Produce json
// @Success 200 {object} ResponseHTTP{}
// @Failure 503 {object} ResponseHTTP{}
// @Router /v1/region [get]
func GetRegion(c *fiber.Ctx) error {
	appF := app.Fiber{C: c}
	// 使用axios进行请求
	queryParams := url.Values{}
	queryParams.Add("subdistrict", "3")
	queryParams.Add("extensions", "base")
	queryParams.Add("key", "22ed02550a66cba6f2340a7072f843e9")
	axiosConfig := &axios.InstanceConfig{}
	axiosConfig.BaseURL = "https://restapi.amap.com"
	resp, err := untils.Request(axiosConfig).Get("/v3/config/district", queryParams)
	if err != nil {
		return appF.Response(fiber.StatusInternalServerError, fiber.StatusInternalServerError, "获取高德省市区", err)
	}
	result := make(map[string]interface{})
	if errJson := json.Unmarshal(resp.Data, &result); errJson != nil {
		return appF.Response(fiber.StatusInternalServerError, fiber.StatusInternalServerError, "解析高德数据失败", errJson)
	}
	return appF.Response(fiber.StatusOK, fiber.StatusOK, "SUCCESS", result)
}
