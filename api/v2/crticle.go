package v2

import (
	"inception-whole/internal/model"
	"inception-whole/internal/model/entity"

	"github.com/gogf/gf/v2/frame/g"
)

type APIArticleListReq struct {
	g.Meta      `path:"/article" method:"get" tags:"博客" summary:"获取博客列表" dc:"获取博客列表"`
	Type        string
	CategoryIDs []int // 栏目ID
	Page        int
	Size        int
}

type APIArticleListRes struct {
	List  []*model.APIContentGetListItem `json:"list" description:"列表"`
	Page  int                            `json:"page" description:"分页码"`
	Size  int                            `json:"size" description:"分页数量"`
	Total int                            `json:"total" description:"数据总数"`
}

type APIArticleDetailReq struct {
	g.Meta `path:"/article/detail" method:"get" tags:"博客" summary:"获取博客列表" dc:"获取博客列表"`
	ID     int
}

type APIArticleDetailRes struct {
	// *model.APIContentGetDetailOutput
	Content  *model.APIContentListItem `json:"blog"`
	Category *model.APICategoryItem    `json:"category"`
	User     *entity.User              `json:"user"`
}
