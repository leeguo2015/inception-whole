package controller

import (
	"context"
	"errors"
	v2 "inception-whole/api/v2"
	"inception-whole/internal/model"
	"inception-whole/internal/service"

	"github.com/gogf/gf/v2/os/glog"
)

// 栏目管理
var APIArticle = APICArticle{}

type APICArticle struct{}

func (a *APICArticle) GetList(ctx context.Context, req *v2.APIArticleListReq) (res *v2.APIArticleListRes, err error) {
	if req.Page < 1 || req.Size < 1 {
		return nil, errors.New("invalid Page or Size, must be greater than 0")
	}
	in := model.APIContentGetListInput{
		Page:        req.Page,
		Size:        req.Size,
		Type:        req.Type,
		CategoryIDs: req.CategoryIDs,
	}
	data, err := service.APIArticle().GetList(ctx, in)
	if err != nil {
		glog.Error(ctx, "GetList failed:", err)
		return nil, err
	}

	res = &v2.APIArticleListRes{
		List:  data.List,
		Page:  data.Page,
		Size:  data.Size,
		Total: data.Total,
	}
	return
}