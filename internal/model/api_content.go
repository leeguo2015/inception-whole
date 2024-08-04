package model

import (
	"inception-whole/internal/model/entity"

	"github.com/gogf/gf/v2/os/gtime"
)

type APIContentGetListItem struct {
	Content  *APIContentListItem `json:"blog"`
	Category *APICategoryItem    `json:"category"`
	User     *entity.User		`json:"user"`
}

type APIContentGetListOutput struct {
	List  []*APIContentGetListItem `json:"list" description:"列表"`
	Page  int                      `json:"page" description:"分页码"`
	Size  int                      `json:"size" description:"分页数量"`
	Total int                      `json:"total" description:"数据总数"`
}

type APIContentGetListInput struct {
	Type        string // 内容模型
	CategoryIDs []int  // 栏目ID
	Page        int    // 分页号码
	Size        int    // 分页数量，最大50
	Sort        int    // 排序类型(0:最新, 默认。1:活跃, 2:热度)
}

type APIContentListItem struct {
	Id         uint        `json:"id"`          // 自增ID
	CategoryId uint        `json:"category_id"` // 栏目ID
	UserId     uint        `json:"user_id"`     // 用户ID
	Title      string      `json:"title"`       // 标题
	Sort       uint        `json:"sort"`        // 排序，数值越低越靠前，默认为添加时的时间戳，可用于置顶
	Brief      string      `json:"brief"`       // 摘要
	Thumb      string      `json:"thumb"`       // 缩略图
	Tags       string      `json:"tags"`        // 标签名称列表，以JSON存储
	Referer    string      `json:"referer"`     // 内容来源，例如github/gitee
	Status     uint        `json:"status"`      // 状态 0: 正常, 1: 禁用
	Content    string      `json:"content"        description:"内容"`
	ViewCount  uint        `json:"view_count"`  // 浏览数量
	ReplyCount uint        `json:"reply_count"` // 回复数量
	ZanCount   uint        `json:"zan_count"`   // 赞
	CaiCount   uint        `json:"cai_count"`   // 踩
	CreatedAt  *gtime.Time `json:"created_at"`  // 创建时间
	UpdatedAt  *gtime.Time `json:"updated_at"`  // 修改时间
}

type APIContentGetDetailInput struct {
	ID int `json:"id"`
}

type APIContentGetDetailOutput struct {
	Content  *APIContentListItem `json:"content"`
	Category *APICategoryItem    `json:"category"`
}
