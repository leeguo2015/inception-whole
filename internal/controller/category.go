package controller

import (
	"context"

	v1 "inception-whole/api/v1"
	"inception-whole/internal/service"
)

// 栏目管理
var Category = cCategory{}

type cCategory struct{}

func (a *cCategory) Tree(ctx context.Context, req *v1.CategoryTreeReq) (res *v1.CategoryTreeRes, err error) {
	res = &v1.CategoryTreeRes{}
	res.List, err = service.Category().GetTree(ctx, req.ContentType)
	return
}
