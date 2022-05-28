package untils

import (
	"io"
	"os"

	"github.com/jinpikaFE/go_fiber/pkg/logging"
)

// 删除文件
func RemoveFile(relatPath string, filename string) bool {
	dir, _ := os.Getwd()
	err := os.Remove(dir + relatPath + "/" + filename)

	if err != nil {
		logging.Error(err)
		return false
	}
	return true

}

// 创建文件夹
func MakeDir(relatPath string) bool {
	dir, _ := os.Getwd()
	err := os.MkdirAll(dir+relatPath, os.ModePerm)
	if err != nil {
		logging.Error(err)
		return false
	}
	return true
}

// 写入文件
func WriteFile(relatPath string, file io.Reader, filename string) bool {
	f, err := os.OpenFile("."+relatPath+"/"+filename, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		logging.Error(err)
		return false
	}
	defer f.Close()

	io.Copy(f, file)
	return true
}
