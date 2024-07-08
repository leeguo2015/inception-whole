package controller

import (
	"context"
	v2 "inception-whole/api/v2"

	"inception-whole/internal/service"

	"github.com/gogf/gf/v2/os/glog"
)

// 栏目管理
var APICategory = APICCategory{}

type APICCategory struct{}

func (a *APICCategory) ALL(ctx context.Context, req *v2.CategoryInceptionListReq) (res *v2.CategoryInceptionListRes, err error) {
	res = &v2.CategoryInceptionListRes{}
	res.List, err = service.APICategory().GetALL(ctx)
	if err != nil {
		glog.Error(ctx, err)
	}
	glog.Debug(ctx, res)
	return
}
