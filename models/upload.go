package models

import (
	"context"
	"fmt"
	"mime/multipart"
	"os"

	"github.com/jinpikaFE/go_fiber/pkg/logging"
	"github.com/jinpikaFE/go_fiber/pkg/tencentcos"
)

func UploadFile(filePath string, fileName string) bool {
	dir, _ := os.Getwd()
	_, _, err := tencentcos.Client.Object.Upload(context.Background(), filePath, fmt.Sprintf("%s/file/%s", dir, fileName), nil)
	if err != nil {
		logging.Error(err)
		return false
	}
	return true
}

func PutFile(filePath string, fileToUpload multipart.File) (bool, error) {
	_, err := tencentcos.Client.Object.Put(context.Background(), filePath, fileToUpload, nil)
	if err != nil {
		logging.Error(err)
		return false, err
	}
	return true, nil
}
