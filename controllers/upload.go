package controller

import (
	"fmt"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/jinpikaFE/go_fiber/pkg/logging"
	"github.com/jinpikaFE/go_fiber/pkg/setting"

	"github.com/jinpikaFE/go_fiber/models"
)

func Upload(c *fiber.Ctx) error {
	file, err := c.FormFile("file")
	message := "SUCCESS"
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

	filePath := "/file/" + fmt.Sprintf("%d", time.Now().Unix()) + "." + strings.Split(file.Filename, `.`)[1]

	if err != nil {
		message = "ERROR"
		logging.Error(err)
	}

	res := models.UploadFile(relFile, filePath)

	if !res {
		message = "上传失败"
	}

	url := setting.CosUrl + filePath

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"code":    "code",
		"message": message,
		"data":    url,
	})
}
