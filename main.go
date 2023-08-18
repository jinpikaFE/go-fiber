package main

import (
	"context"
	"fmt"

	_ "github.com/jinpikaFE/go_fiber/docs"
	"github.com/jinpikaFE/go_fiber/database"
	"github.com/jinpikaFE/go_fiber/pkg/gredis"
	"github.com/jinpikaFE/go_fiber/pkg/logging"
	"github.com/jinpikaFE/go_fiber/pkg/setting"
	"github.com/jinpikaFE/go_fiber/pkg/tencent"
	"github.com/jinpikaFE/go_fiber/routers"
)

func init() {
	gredis.Setup()
	tencent.SetupSms()
}

func main() {
	app := routers.InitRouter()
	// 确保在main函数结束后关闭连接
	defer func() {
		if err := database.MgbClient.Disconnect(context.Background()); err != nil {
			logging.Fatal(err)
		}
	}()

	app.Listen(fmt.Sprintf(":%d", setting.HTTPPort))
}
