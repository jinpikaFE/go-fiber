package e

const (
	CACHE_ARTICLE = "ARTICLE"
	CACHE_TAG     = "TAG"
)

type PageStruct struct {
	Page     int `validate:"required" query:"page" json:"page" xml:"page" form:"page"`
	PageSize int `validate:"required" query:"pageSize" json:"pageSize" xml:"pageSize" form:"pageSize"`
}
