package models

import (
	"context"
	"io"

	"github.com/jinpikaFE/go_fiber/pkg/logging"
	"github.com/jinpikaFE/go_fiber/pkg/tencentcos"
)

func UploadFile(file io.Reader, filePath string) bool {
	_, err := tencentcos.Client.Object.Put(context.Background(), filePath, file, nil)
	if err != nil {
		logging.Error(err)
		return false
	}
	return true
}
