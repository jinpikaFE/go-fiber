package main

import (
	"fmt"

	"github.com/jinpikaFE/go_fiber/pkg/gredis"
	"github.com/jinpikaFE/go_fiber/pkg/setting"
	"github.com/jinpikaFE/go_fiber/routers"
)

func init() {
	gredis.Setup()
}

func main() {
	app := routers.InitRouter()

	app.Listen(fmt.Sprintf(":%d", setting.HTTPPort))
}
