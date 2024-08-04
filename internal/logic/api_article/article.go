package api_article

import (
	"context"

	"inception-whole/internal/dao"
	"inception-whole/internal/logic/api_category"
	"inception-whole/internal/model"
	"inception-whole/internal/model/entity"
	"inception-whole/internal/service"

	"github.com/gogf/gf/v2/os/glog"
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
	if in.Size == 0 {
		in.Size = 10
	}
	out = &model.APIContentGetListOutput{
		Page: in.Page,
		Size: in.Size,
	}
	if len(in.CategoryIDs) > 0 {
		m = m.WhereIn(dao.Content.Columns().CategoryId, in.CategoryIDs)
	}
	if in.Type != "" {
		m = m.Where(dao.Content.Columns().Type, in.Type)
	}
	m.Fields("LEFT(content, 100) AS content")

	listModel := m.Page(in.Page, in.Size)
	out.Total, err = m.Count()
	if err != nil {
		return out, err
	}
	// Content
	if err := listModel.ScanList(&out.List, "Content"); err != nil {
		return out, err
	}
	categories, err := api_category.New().GetList(ctx, "")
	if err != nil {
		return out, err
	}
	for i := 0; i < len(out.List); i++ {
		for _, category := range categories {
			if category.Id == out.List[i].Content.CategoryId {
				out.List[i].Category = &model.APICategoryItem{
					Id:      category.Id,
					Name:    category.Name,
					Content: category.Content,
					Brief:   category.Brief,
				}
				break
			}
		}
		// user :=
		glog.Debug(ctx, "out.List[i].Content.UserId", out.List[i].Content.UserId)
		u := &entity.User{}
		if err := dao.User.Ctx(ctx).Where(dao.User.Columns().Id, out.List[i].Content.UserId).Scan(u); err != nil {
			glog.Error(ctx, err)
		}
		out.List[i].User = u

	}

	return
}

func (s *APISAriticle) Detail(ctx context.Context, in model.APIContentGetDetailInput) (out *model.APIContentGetDetailOutput, err error) {
	glog.Debug(ctx, in.ID)

	var (
		m = dao.Content.Ctx(ctx)
	)
	out = &model.APIContentGetDetailOutput{
		Content:  &model.APIContentListItem{},
		Category: &model.APICategoryItem{},
	}
	m = m.Where(dao.Content.Columns().Id, in.ID)
	if err := m.Scan(&out.Content); err != nil {
		return out, err
	}
	glog.Debug(ctx, in.ID)
	glog.Debug(ctx, "out.Content", out.Content)

	// err = m.Scan(out.Content)
	if err != nil {
		return nil, err
	}
	return
}
