package database

import (
	"context"

	"github.com/jinpikaFE/go_fiber/pkg/logging"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MgbClient *mongo.Client

/** 数据库驱动 链接 monitor-fiber */
var MgbDatabase *mongo.Database

func init() {
	var err error
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	// 连接到MongoDB
	MgbClient, err = mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		logging.Fatal(err)
	}
	MgbDatabase = MgbClient.Database("monitor-fiber")
}
