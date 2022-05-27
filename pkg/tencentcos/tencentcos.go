package tencentcos

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/jinpikaFE/go_fiber/pkg/setting"
	"github.com/tencentyun/cos-go-sdk-v5"
)

var (
	Client *cos.Client
)

func init() {
	u, _ := url.Parse(setting.CosUrl)
	b := &cos.BaseURL{BucketURL: u}
	Client = cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  fmt.Sprintf("%s", setting.SecretId),  // 替换为用户的 SecretId，请登录访问管理控制台进行查看和管理，https://console.cloud.tencent.com/cam/capi
			SecretKey: fmt.Sprintf("%s", setting.SecretKey), // 替换为用户的 SecretKey，请登录访问管理控制台进行查看和管理，https://console.cloud.tencent.com/cam/capi
		},
	})
}
