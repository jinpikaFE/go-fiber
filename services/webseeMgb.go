package services

import (
	"context"
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/jinpikaFE/go_fiber/database"
	"github.com/jinpikaFE/go_fiber/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func SetMgbData(data map[string]interface{}) (*mongo.InsertOneResult, error) {
	if data["message"] != nil {
		var jsonValue interface{}
		jsonData := data["message"].(string)
		err := json.Unmarshal([]byte(jsonData), &jsonValue)
		if err != nil {
			fmt.Println("解析时间错误:", err)
			data["message"] = jsonData

		} else {
			data["message"] = jsonValue
		}

	}

	if data["type"] == "xhr" || data["type"] == "fetch" {
		data["type"] = "apiErr"
	}

	res, err := database.MgbDatabase.Collection(data["type"].(string)).InsertOne(context.Background(), data)
	return res, err
}

func GetMgbMonitor(pageNum int, pageSize int, maps *models.MgbMonitorCondition, inequality *models.MgbMonitorInequality) (interface{}, interface{}, error) {
	timeLayout := "2006-01-02 15:04:05"
	// 解析时间字符串
	parsedTimeStart, err1 := time.Parse(timeLayout, inequality.StartTime)
	if err1 != nil {
		fmt.Println("解析时间错误:", err1.Error())
		return nil, 0, err1
	}
	// 转换为时间戳
	timestampStart := parsedTimeStart.UnixMilli()
	parsedTimeEnd, err2 := time.Parse(timeLayout, inequality.EndTime)
	if err2 != nil {
		fmt.Println("解析时间错误:", err1.Error())
		return nil, 0, err2
	}
	// 转换为时间戳
	timestampEnd := parsedTimeEnd.UnixMilli()

	query := buildQueryFromStruct(maps)

	fmt.Print(timestampStart, timestampEnd)

	query["time"] = bson.M{
		"$gte": timestampStart,
		"$lte": timestampEnd,
	}

	findOptions := options.Find()
	findOptions.SetSort(bson.M{"time": -1})
	findOptions.SetSkip(int64(pageNum))
	findOptions.SetLimit(int64(pageSize))

	// 执行查询获取分页数据
	cur, err := database.MgbDatabase.Collection(maps.Type).Find(context.Background(), query, findOptions)
	if err != nil {
		return nil, nil, err
	}
	defer cur.Close(context.Background())
	var dataList []map[string]interface{}

	for cur.Next(context.Background()) {
		var data map[string]interface{}
		err := cur.Decode(&data)
		if err != nil {
			return nil, nil, err
		}
		dataList = append(dataList, data)
	}

	// 获取总记录数
	totalCount, err := database.MgbDatabase.Collection(maps.Type).CountDocuments(context.Background(), query)
	if err != nil {
		return nil, nil, err
	}
	return dataList, totalCount, err
}

/** 构建反射 结构体 */
func buildQueryFromStruct(params interface{}) bson.M {
	query := bson.M{}

	paramsValue := reflect.Indirect(reflect.ValueOf(params))
	paramsType := paramsValue.Type()

	for i := 0; i < paramsType.NumField(); i++ {
		field := paramsType.Field(i)
		tag := field.Tag.Get("query")
		value := paramsValue.Field(i).String()

		if value != "" {
			query[strings.ToLower(tag)] = value
		}
	}

	return query
}

func GetMgbRecordScreen(recordScreenId string) (map[string]interface{}, error) {

	var result map[string]interface{}

	err2 := database.MgbDatabase.Collection("recordScreen").FindOne(context.Background(), bson.M{
		"recordScreenId": recordScreenId,
	}).Decode(&result)
	if err2 != nil {
		return nil, err2
	}
	// 返回 JSON
	return result, nil
}

// 各项指标数据统计
func GetStatistRes(params *models.MgbStatistDataParamsInequality) (*models.MgbStatistData, error) {
	res := &models.MgbStatistData{}
	query := bson.M{}
	if params.StartTime != "" && params.EndTime != "" {
		timeLayout := "2006-01-02 15:04:05"
		// 解析时间字符串
		parsedTimeStart, err1 := time.Parse(timeLayout, params.StartTime)
		if err1 != nil {
			fmt.Println("解析时间错误:", err1.Error())
			return nil, err1
		}
		// 转换为时间戳
		timestampStart := parsedTimeStart.UnixMilli()
		parsedTimeEnd, err2 := time.Parse(timeLayout, params.EndTime)
		if err2 != nil {
			fmt.Println("解析时间错误:", err1.Error())
			return nil, err2
		}
		// 转换为时间戳
		timestampEnd := parsedTimeEnd.UnixMilli()

		query["time"] = bson.M{
			"$gte": timestampStart,
			"$lte": timestampEnd,
		}
	}

	// 获取uv总记录数
	uvCount, errUv := database.MgbDatabase.Collection("uv").CountDocuments(context.Background(), query)
	if errUv != nil {
		return nil, errUv
	}
	res.UVNum = uvCount

	// 获取pv总记录数
	pvCount, errPv := database.MgbDatabase.Collection("pv").CountDocuments(context.Background(), query)
	if errPv != nil {
		return nil, errPv
	}
	res.PVNum = pvCount

	return res, nil
}
