package untils

import (
	"encoding/json"
	"regexp"

	"github.com/influxdata/influxdb-client-go/v2/api"
)

//mobile verify
func VerifyMobileFormat(mobileNum string) bool {
	regular := "^((13[0-9])|(14[5,7])|(15[0-3,5-9])|(17[0,3,5-8])|(18[0-9])|166|198|199|(147))\\d{8}$"

	reg := regexp.MustCompile(regular)
	return reg.MatchString(mobileNum)
}

// influxdb 查询结果处理
func InfluxdbQueryResult(result *api.QueryTableResult) ([]map[string]interface{}, error) {
	results := []map[string]interface{}{}
	for result.Next() {
		record := result.Record()
		item := map[string]interface{}{}
		for key, value := range record.Values() {
			itemValue := value
			if v, ok := itemValue.(string); ok {
				// 如果当前值是字符串，则尝试解析其为 JSON 对象、数组或数值
				var jsonValue interface{}
				err := json.Unmarshal([]byte(v), &jsonValue)
				if err == nil {
					switch realValue := jsonValue.(type) {
					case map[string]interface{}:
						// 当前值是 JSON 对象
						itemValue = realValue
					case []interface{}:
						// 当前值是 JSON 数组
						itemValue = realValue
					default:
						// 当前值是 JSON 数值、字符串、布尔值或 null
						itemValue = realValue
					}
				}
			}
			item[key] = itemValue
		}
		results = append(results, item)
	}
	if err := result.Err(); err != nil {
		return nil, err
	}
	return results, nil
}

// 结构转map
func StructToMap(obj interface{}) map[string]interface{} {
	// 将结构体转换成字节数组
	b, err := json.Marshal(obj)
	if err != nil {
		return nil
	}
	// 解码字节数组为 map
	var result map[string]interface{}
	if err := json.Unmarshal(b, &result); err != nil {
		return nil
	}
	return result
}
