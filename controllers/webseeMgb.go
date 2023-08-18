package controller

import (
	"encoding/json"
	"encoding/xml"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/jinpikaFE/go_fiber/models"
	"github.com/jinpikaFE/go_fiber/pkg/app"
	"github.com/jinpikaFE/go_fiber/pkg/e"
	"github.com/jinpikaFE/go_fiber/pkg/logging"
	"github.com/jinpikaFE/go_fiber/pkg/valodates"
	"github.com/jinpikaFE/go_fiber/services"
)

// 添加Monodb监控数据
// @Summary 添加Monodb监控数据
// @Description 添加Monodb监控数据
// @Tags Monodb监控数据处理
// @Accept json
// @Produce json
// @Success 200 {object} ResponseHTTP{}
// @Failure 503 {object} ResponseHTTP{}
// @Router /v1/mgb/monitor [post]
func SetMgbData(c *fiber.Ctx) error {
	appF := app.Fiber{C: c}
	// data := &models.ReportData{}
	var data map[string]interface{}

	switch ct := c.Get("Content-Type"); ct {
	case "application/json":
		if err := json.Unmarshal(c.Body(), &data); err != nil {
			return err
		}
	case "text/plain;charset=UTF-8":
		if err := json.Unmarshal([]byte(c.Body()), &data); err != nil {
			return err
		}
	case "application/xml":
		if err := xml.Unmarshal(c.Body(), &data); err != nil {
			return err
		}
	default:
		return appF.Response(fiber.StatusInternalServerError, fiber.StatusInternalServerError, "参数解析错误", fmt.Errorf("unsupported Content-Type: %v", ct))
	}
	p, err := services.SetMgbData(data)
	if err != nil {
		return appF.Response(fiber.StatusInternalServerError, fiber.StatusInternalServerError, "ERROR", err)
	}
	return appF.Response(fiber.StatusOK, fiber.StatusOK, "SUCCESS", p)
}

// 获取Monodb监控数据
// @Summary 获取Monodb监控数据
// @Description 获取Monodb监控数据
// @Tags Monodb监控数据处理
// @Accept json
// @Produce json
// @Success 200 {object} ResponseHTTP{}
// @Failure 503 {object} ResponseHTTP{}
// @Router /v1/mgb/monitor [get]
func GetMgbMonitor(c *fiber.Ctx) error {
	appF := app.Fiber{C: c}
	page := &e.PageStruct{}
	maps := &models.MgbMonitorCondition{}
	inequality := &models.MgbMonitorInequality{}
	err2 := c.QueryParser(page)
	err3 := c.QueryParser(inequality)

	if err := c.QueryParser(maps); err != nil && err2 != nil && err3 != nil {
		logging.Error(err)
		return appF.Response(fiber.StatusInternalServerError, fiber.StatusInternalServerError, "参数解析错误", err)
	}

	// 入参验证
	if errors := valodates.ValidateStruct(*page); errors != nil {
		return appF.Response(fiber.StatusBadRequest, fiber.StatusBadRequest, "检验参数错误", errors)
	}

	// 入参验证
	if errors := valodates.ValidateStruct(*maps); errors != nil {
		return appF.Response(fiber.StatusBadRequest, fiber.StatusBadRequest, "检验参数错误", errors)
	}

	// 入参验证
	if errors := valodates.ValidateStruct(*inequality); errors != nil {
		return appF.Response(fiber.StatusBadRequest, fiber.StatusBadRequest, "检验参数错误", errors)
	}

	p, resultCount, err := services.GetMgbMonitor((page.Page-1)*page.PageSize, page.PageSize, maps, inequality)
	if err != nil {
		return appF.Response(fiber.StatusInternalServerError, fiber.StatusInternalServerError, "查询失败", err)
	}
	result := make(map[string]interface{})
	result["list"] = p
	result["pageNum"] = page.Page
	result["pageSize"] = page.PageSize
	result["total"] = resultCount
	return appF.Response(fiber.StatusOK, fiber.StatusOK, "SUCCESS", result)
}

// 获取Monodb录屏数据
// @Summary 获取Monodb录屏数据
// @Description 获取Monodb录屏数据
// @Tags Monodb监控数据处理
// @Accept json
// @Produce json
// @Success 200 {object} ResponseHTTP{}
// @Failure 503 {object} ResponseHTTP{}
// @Router /v1/mgb/monitor/screen/:id [get]
func GetMgbRecordScreen(c *fiber.Ctx) error {
	appF := app.Fiber{C: c}
	id := c.Params("id")

	p, err := services.GetMgbRecordScreen(id)
	if err != nil {
		return appF.Response(fiber.StatusInternalServerError, fiber.StatusInternalServerError, "查询失败", err)
	}
	return appF.Response(fiber.StatusOK, fiber.StatusOK, "SUCCESS", p)
}

// 获取统计数据
// @Summary 获取统计数据
// @Description 获取统计数据
// @Tags Monodb监控数据处理
// @Accept json
// @Produce json
// @Success 200 {object} ResponseHTTP{}
// @Failure 503 {object} ResponseHTTP{}
// @Router /v1/mgb/monitor/statist [get]
func GetStatistRes(c *fiber.Ctx) error {
	appF := app.Fiber{C: c}
	inequality := &models.MgbStatistDataParamsInequality{}
	if err := c.QueryParser(inequality); err != nil {
		logging.Error(err)
		return appF.Response(fiber.StatusInternalServerError, fiber.StatusInternalServerError, "参数解析错误", err)
	}

	p, err := services.GetStatistRes(inequality)
	if err != nil {
		return appF.Response(fiber.StatusInternalServerError, fiber.StatusInternalServerError, "查询失败", err)
	}
	return appF.Response(fiber.StatusOK, fiber.StatusOK, "SUCCESS", p)
}
