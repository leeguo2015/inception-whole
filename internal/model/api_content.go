package model

type APIContentGetListItem struct {
	Content  *ContentListItem         `json:"content"`
	Category *ContentListCategoryItem `json:"category"`
}

type APIContentGetListOutput struct {
	List  []*APIContentGetListItem `json:"list" description:"列表"`
	Page  int                     `json:"page" description:"分页码"`
	Size  int                     `json:"size" description:"分页数量"`
	Total int                     `json:"total" description:"数据总数"`
}

type APIContentGetListInput struct {
	Type        string // 内容模型
	CategoryIDs []int  // 栏目ID
	Page        int    // 分页号码
	Size        int    // 分页数量，最大50
	Sort        int    // 排序类型(0:最新, 默认。1:活跃, 2:热度)
}
