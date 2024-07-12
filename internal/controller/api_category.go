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

func (a *APICCategory) ALL(ctx context.Context, req *v2.APICategoryListReq) (res *v2.APICategoryListRes, err error) {
	res = &v2.APICategoryListRes{}
	res.List, err = service.APICategory().GetALL(ctx)
	if err != nil {
		glog.Error(ctx, err)
	}
	return
}
