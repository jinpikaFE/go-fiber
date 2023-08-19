package database

import (
	"context"
	"fmt"

	"github.com/jinpikaFE/go_fiber/pkg/logging"
	"github.com/jinpikaFE/go_fiber/pkg/setting"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MgbClient *mongo.Client

/** 数据库驱动 链接 monitor-fiber */
var MgbDatabase *mongo.Database

func init() {

	var err error
	clientOptions := options.Client().ApplyURI(fmt.Sprintf("mongodb://%s:%s@%s", setting.MongoSetting.UserName, setting.MongoSetting.Password, setting.MongoSetting.Host))

	// 连接到MongoDB
	MgbClient, err = mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		logging.Fatal(err)
	}
	MgbDatabase = MgbClient.Database("monitor-fiber")
}
