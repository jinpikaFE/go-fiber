package controller

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/jinpikaFE/go_fiber/pkg/app"
	"github.com/jinpikaFE/go_fiber/pkg/logging"
	"github.com/jinpikaFE/go_fiber/pkg/setting"
	"github.com/jinpikaFE/go_fiber/pkg/untils"

	"github.com/jinpikaFE/go_fiber/models"
)

func Upload(c *fiber.Ctx) error {
	appF := app.Fiber{C: c}
	logging.Info("/v1/upload")
	file, err := c.FormFile("file")
	logging.Info(*file)
	if err != nil {
		logging.Error((err))
		return appF.Response(fiber.StatusBadRequest, fiber.StatusBadRequest, "file为空", nil)
	}

	// //大小限制2Mb
	// if file.Size > (2 << 20) {
	// 	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
	// 		"code":    fiber.StatusBadRequest,
	// 		"message": "文件过大 超过2m",
	// 		"data":    nil,
	// 	})
	// }

	relFile, err := file.Open()

	relatPath := "/file"
	if res := untils.MakeDir(relatPath); !res {
		return appF.Response(fiber.StatusInternalServerError, fiber.StatusInternalServerError, "创建文件夹失败", nil)
	}

	if res := untils.WriteFile(relatPath, relFile, file.Filename); !res {
		return appF.Response(fiber.StatusInternalServerError, fiber.StatusInternalServerError, "写入文件失败", nil)
	}

	filePath := "/file/" + fmt.Sprintf("%s", time.Now().Format("2006-01-02")) + "/" + file.Filename

	if err != nil {
		logging.Error(err)
		return appF.Response(fiber.StatusInternalServerError, fiber.StatusInternalServerError, "上传cos失败", err)
	}

	res := models.UploadFile(filePath, file.Filename)

	if res := untils.RemoveFile(relatPath, file.Filename); !res {
		return appF.Response(fiber.StatusInternalServerError, fiber.StatusInternalServerError, "删除本地文件失败", err)
	}

	if !res {
		return appF.Response(fiber.StatusInternalServerError, fiber.StatusInternalServerError, "上传失败", err)
	}

	url := setting.CosUrl + filePath

	data := map[string]interface{}{"url": url}

	return appF.Response(fiber.StatusOK, fiber.StatusOK, "SUCCESS", data)
}
