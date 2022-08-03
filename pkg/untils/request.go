package untils

import (
	"github.com/vicanso/go-axios"
)

func Request(axiosConfig *axios.InstanceConfig) *axios.Instance {
	request := axios.NewInstance(&axios.InstanceConfig{
		BaseURL:     axiosConfig.BaseURL,
		EnableTrace: true,
	})

	return request
}
