package v2

import (
	"inception-whole/internal/model"

	"github.com/gogf/gf/v2/frame/g"
)

type CategoryInceptionListReq struct {
	g.Meta `path:"/category" method:"get" tags:"分类" summary:"获取分类列表InceptionList" dc:"获取所有分类列表"`
}

type CategoryInceptionListRes struct {
	List []*model.CategoryInceptionItem `json:"list" dc:"栏目列表"`
}
