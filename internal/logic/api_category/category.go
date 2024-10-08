package api_category

import (
	"context"
	"time"

	"inception-whole/internal/service"

	"github.com/gogf/gf/v2/os/gcache"

	"inception-whole/internal/dao"
	"inception-whole/internal/model"
	"inception-whole/internal/model/entity"
)

type APISCategory struct{}

const (
	inceptionCacheDuration = time.Hour
	inceptionCacheKey      = "incepation_APISCategory_cache"
)

func init() {
	service.RegisterAPICCategory(New())
}

func New() *APISCategory {
	return &APISCategory{}
}

// GetList 获得所有的栏目列表。
func (s *APISCategory) GetList(ctx context.Context, Type string) (list []*entity.Category, err error) {
	d := dao.Category.Ctx(ctx).
		OrderAsc(dao.Category.Columns().Sort).
		OrderAsc(dao.Category.Columns().Id)
	if Type != "" {
		d = d.Where(dao.Category.Columns().ContentType, Type)
	}
	err = d.Scan(&list)
	return
}

// GetTree 查询列表
func (s *APISCategory) GetALL(ctx context.Context, Type string) ([]*model.APICategoryItem, error) {
	// 缓存控制
	var (
		cacheKeyT = inceptionCacheKey + Type
		cacheFunc = func(ctx context.Context) (interface{}, error) {
			entities, err := s.GetList(ctx, Type)
			if err != nil {
				return nil, err
			}
			tree, err := s.formIncepation(entities)
			if err != nil {
				return nil, err
			}
			return tree, nil
		}
	)
	v, err := gcache.GetOrSetFunc(ctx, cacheKeyT, cacheFunc, inceptionCacheDuration)
	if err != nil {
		return nil, err
	}
	var (
		result []*model.APICategoryItem
	)
	err = v.Scan(&result)
	return result, err
}

func (s *APISCategory) formIncepation(entities []*entity.Category) ([]*model.APICategoryItem, error) {
	tree := make([]*model.APICategoryItem, 0)
	for _, entity := range entities {
		tree = append(tree, &model.APICategoryItem{
			Id:      entity.Id,
			Name:    entity.Name,
			Content: entity.Brief,
		})

	}
	return tree, nil
}
