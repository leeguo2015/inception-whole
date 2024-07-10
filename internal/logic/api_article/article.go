package api_aritcle

import (
	"context"

	"inception-whole/internal/dao"
	"inception-whole/internal/model"
	"inception-whole/internal/model/entity"
	"inception-whole/internal/service"
)

type APISAriticle struct{}

const ()

func init() {
	service.RegisterAPIIArticle(New())
}

func New() *APISAriticle {
	return &APISAriticle{}
}

func (s *APISAriticle) GetList(ctx context.Context, in model.APIContentGetListInput) (out *model.APIContentGetListOutput, err error) {
	var (
		m = dao.Content.Ctx(ctx)
	)

	out = &model.APIContentGetListOutput{
		Page: in.Page,
		Size: in.Size,
	}

	if in.Type != "" {
		m = m.Where(dao.Content.Columns().Type, in.Type)
	}
	if len(in.CategoryIDs) > 0 {
		m = m.WhereIn(dao.Content.Columns().CategoryId, in.CategoryIDs)
	}
	if in.Type != "" {
		m = m.Where(dao.Content.Columns().Type, in.Type)
	}
	listModel := m.Page(in.Page, in.Size)

	// 执行查询
	var list []*entity.Content
	if err := listModel.Scan(&list); err != nil {
		return out, err
	}
	// 没有数据
	if len(list) == 0 {
		return out, nil
	}
	out.Total, err = m.Count()
	if err != nil {
		return out, err
	}
	// Content
	if err := listModel.ScanList(&out.List, "Content"); err != nil {
		return out, err
	}
	return nil, nil
}
