package e

const (
	CACHE_ARTICLE = "ARTICLE"
	CACHE_TAG     = "TAG"
)

type PageStruct struct {
	Page     int `validate:"required" query:"pageNum" json:"pageNum" xml:"pageNum" form:"pageNum"`
	PageSize int `validate:"required" query:"pageSize" json:"pageSize" xml:"pageSize" form:"pageSize"`
}
