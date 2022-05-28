package controller

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/jinpikaFE/go_fiber/pkg/logging"
	"github.com/jinpikaFE/go_fiber/pkg/setting"
	"github.com/jinpikaFE/go_fiber/pkg/untils"

	"github.com/jinpikaFE/go_fiber/models"
)

func Upload(c *fiber.Ctx) error {
	logging.Info("/v1/upload")
	file, err := c.FormFile("file")
	message := "SUCCESS"
	code := 200
	if err != nil {
		logging.Error((err))
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"code":    fiber.StatusBadRequest,
			"message": "file为空",
			"data":    nil,
		})
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
		message = "创建文件夹失败"
		code = fiber.StatusInternalServerError
	}

	if res := untils.WriteFile(relatPath, relFile, file.Filename); !res {
		message = "写入文件失败"
		code = fiber.StatusInternalServerError
	}

	filePath := "/file/" + fmt.Sprintf("%s", time.Now().Format("2006-01-02")) + "/" + file.Filename

	if err != nil {
		message = "ERROR"
		code = fiber.StatusInternalServerError
		logging.Error(err)
	}

	res := models.UploadFile(filePath, file.Filename)

	if res := untils.RemoveFile(relatPath, file.Filename); !res {
		message = "删除本地文件失败"
		code = fiber.StatusInternalServerError
	}

	if !res {
		code = fiber.StatusInternalServerError
		message = "上传失败"
	}

	url := setting.CosUrl + filePath

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"code":    code,
		"message": message,
		"data":    url,
	})
}
