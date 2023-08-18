package models

type HttpData struct {
	Type         *string `json:"type"`
	Method       *string `json:"method"`
	Time         int64   `json:"time"`
	Url          string  `json:"url"`          // 接口地址
	ElapsedTime  int64   `json:"elapsedTime"`  // 接口时长
	Message      string  `json:"message"`      // 接口信息
	Status       *int64  `json:"status"`       // 接口状态编码
	StatusString *string `json:"statusString"` // 接口状态
	RequestData  struct {
		HttpType string      `json:"httpType"` // 请求类型 xhr fetch
		Method   string      `json:"method"`   // 请求方式
		Data     interface{} `json:"data,omitempty"`
	} `json:"requestData,omitempty"`
	Response struct {
		Status *int64      `json:"status"` // 接口状态
		Data   interface{} `json:"data,omitempty"`
	} `json:"response,omitempty"`
}

type ResourceError struct {
	Time    int64  `json:"time"`
	Message string `json:"message"`
	Name    string `json:"name"`
}

type Attribution struct {
	Name          string `json:"name"`
	EntryType     string `json:"entryType"`
	StartTime     int    `json:"startTime"`
	Duration      int    `json:"duration"`
	ContainerType string `json:"containerType"`
	ContainerSrc  string `json:"containerSrc"`
	ContainerID   string `json:"containerId"`
	ContainerName string `json:"containerName"`
}

type LongTask struct {
	StartTime   float32        `json:"startTime"`
	EntryType   string         `json:"entryType"`
	Duration    int            `json:"duration"`
	Name        string         `json:"name"`
	Attribution []*Attribution `json:"attribution"`
}

type PerformanceData struct {
	Name   string `json:"name"`
	Value  int64  `json:"value"`
	Rating string `json:"rating"`
}

type Memory struct {
	JSHeapSizeLimit int64 `json:"jsHeapSizeLimit"`
	TotalJSHeapSize int64 `json:"totalJSHeapSize"`
	UsedJSHeapSize  int64 `json:"usedJSHeapSize"`
}

type CodeError struct {
	Column   int64  `json:"column"`
	Line     int64  `json:"line"`
	Message  string `json:"message"`
	FileName string `json:"fileName"`
}

type Behavior struct {
	Type     string      `json:"type"`
	Category interface{} `json:"category"`
	Status   int64       `json:"status"`
	Time     int64       `json:"time"`
	Data     interface{} `json:"data"`
	Message  string      `json:"message"`
	Name     *string     `json:"name,omitempty"`
}

type RecordScreen struct {
	RecordScreenId string `json:"recordScreenId"`
	Events         string `json:"events"`
}

type DeviceInfo struct {
	BrowserVersion interface{} `json:"browserVersion"` // 版本号
	Browser        string      `json:"browser"`        // Chrome
	OSVersion      interface{} `json:"osVersion"`      // 电脑系统 10
	OS             string      `json:"os"`             // 设备系统
	UA             string      `json:"ua"`             // 设备详情
	Device         string      `json:"device"`         // 设备种类描述
	DeviceType     string      `json:"device_type"`    // 设备种类，如pc
}

type RequestData struct {
	HttpType string      `json:"httpType"`
	Method   string      `json:"method"`
	Data     interface{} `json:"data"`
}

type Response struct {
	Status int         `json:"status"`
	Data   interface{} `json:"data,omitempty"`
}

type ReportData struct {
	Name            string            `json:"name"`
	Type            string            `json:"type"`
	PageUrl         string            `json:"pageUrl"`
	Rating          string            `json:"rating" description:"性能指标 poor or good"`
	Time            int64             `json:"time"`
	UUID            string            `json:"uuid"`
	Apikey          string            `json:"apikey"`
	Status          string            `json:"status"`
	SdkVersion      string            `json:"sdkVersion"`
	Events          string            `json:"events"`
	UserId          string            `json:"userId"`
	Line            int32             `json:"line" description:"发生错误位置"`
	Column          int32             `json:"column" description:"发生错误位置"`
	Message         string            `json:"message" description:"相关信息"`
	RecordScreenId  string            `json:"recordScreenId" description:"录屏信息id"`
	FileName        string            `json:"fileName" description:"错误文件"`
	Url             string            `json:"url" description:"请求url"`
	ElapsedTime     int32             `json:"elapsedTime" description:"接口时长"` // 接口时长
	Value           float32           `json:"value,omitempty" description:"性能指标值"`
	RequestData     *RequestData      `json:"requestData,omitempty" description:"请求数据"`
	Response        *Response         `json:"response,omitempty" description:"响应数据"`
	Breadcrumb      []*BreadcrumbData `json:"breadcrumb,omitempty" description:"用户行为"`
	HttpData        *HttpData         `json:"httpData,omitempty"`
	ResourceError   *ResourceError    `json:"resourceError,omitempty"`
	LongTask        *LongTask         `json:"longTask,omitempty"`
	PerformanceData *PerformanceData  `json:"performanceData,omitempty"`
	Memory          *Memory           `json:"memory,omitempty"`
	CodeError       *CodeError        `json:"codeError,omitempty"`
	RecordScreen    *RecordScreen     `json:"recordScreen,omitempty"`
	DeviceInfo      *DeviceInfo       `json:"deviceInfo,omitempty"`
}

type ReportDataJson struct {
	Name            string  `json:"name"`
	Type            string  `json:"type"`
	PageUrl         string  `json:"pageUrl"`
	Rating          string  `json:"rating" description:"性能指标 poor or good"`
	Time            int64   `json:"time"`
	UUID            string  `json:"uuid"`
	Apikey          string  `json:"apikey"`
	Status          string  `json:"status"`
	SdkVersion      string  `json:"sdkVersion"`
	Events          string  `json:"events"`
	UserId          string  `json:"userId"`
	Line            int32   `json:"line" description:"发生错误位置"`
	Column          int32   `json:"column" description:"发生错误位置"`
	Message         string  `json:"message" description:"相关信息"`
	RecordScreenId  string  `json:"recordScreenId"`
	FileName        string  `json:"fileName" description:"错误文件"`
	Url             string  `json:"url" description:"请求url"`
	ElapsedTime     int32   `json:"elapsedTime" description:"接口时长"` // 接口时长
	Value           float32 `json:"value,omitempty" description:"性能指标值"`
	RequestData     string  `json:"requestData,omitempty" description:"请求数据"`
	Response        string  `json:"response,omitempty" description:"响应数据"`
	Breadcrumb      string  `json:"breadcrumb,omitempty" description:"用户行为"`
	HttpData        string  `json:"httpData,omitempty"`
	ResourceError   string  `json:"resourceError,omitempty"`
	LongTask        string  `json:"longTask,omitempty"`
	PerformanceData string  `json:"performanceData,omitempty"`
	Memory          string  `json:"memory,omitempty"`
	CodeError       string  `json:"codeError,omitempty"`
	RecordScreen    string  `json:"recordScreen,omitempty"`
	DeviceInfo      string  `json:"deviceInfo,omitempty"`
}

type ResourceTarget struct {
	Src       *string `json:"src,omitempty"`
	Href      *string `json:"href,omitempty"`
	LocalName *string `json:"localName,omitempty"`
}

type AuthInfo struct {
	Apikey     string  `json:"apiKey"`
	SdkVersion string  `json:"sdkVersion"`
	UserId     *string `json:"userId,omitempty"`
}

type BreadcrumbData struct {
	Type     string      `json:"type" description:"事件类型"`
	Category string      `json:"category" description:"用户行为类型"`
	Status   string      `json:"status" description:"行为状态"`
	Time     int64       `json:"time"`
	Data     interface{} `json:"data,omitempty"`
}

type MonitorParams struct {
	UUID           string `query:"uuid" json:"uuid" xml:"uuid" form:"uuid"`
	UserId         string `query:"userId" json:"userId" xml:"userId" form:"userId"`
	RecordScreenId string `query:"recordScreenId" json:"recordScreenId" xml:"recordScreenId" form:"recordScreenId"`
	// 项目名，项目key
	Apikey         string `query:"apikey" json:"apikey" xml:"apikey" form:"apikey"`
	Name           string `query:"name" json:"name" xml:"name" form:"name"`
	// 类型对应不同数据表或者是集合
	Type           string `validate:"required" query:"type" json:"type" xml:"type" form:"type"`
	// query tag是query参数别名，json xml，form适合post // validate:"required" 
	StartTime string `validate:"required" query:"startTime" json:"startTime" xml:"startTime" form:"startTime"`
	EndTime   string `validate:"required" query:"endTime" json:"endTime" xml:"endTime" form:"endTime"`
}
