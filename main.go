package main

import (
	"fmt"

	"github.com/jinpikaFE/go_fiber/pkg/gredis"
	"github.com/jinpikaFE/go_fiber/pkg/setting"
	"github.com/jinpikaFE/go_fiber/pkg/tencent"
	"github.com/jinpikaFE/go_fiber/routers"
	_ "github.com/jinpikaFE/go_fiber/docs"
)

func init() {
	gredis.Setup()
	tencent.SetupSms()
}

func main() {
	app := routers.InitRouter()

	app.Listen(fmt.Sprintf(":%d", setting.HTTPPort))
}
