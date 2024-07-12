
package service

import (
	"context"
	"inception-whole/internal/model"
	// "inception-whole/internal/model"
	// "inception-whole/internal/model/entity"
)

type APIIArticle interface {
	GetList(ctx context.Context, in model.APIContentGetListInput) (out *model.APIContentGetListOutput, err error)
	Detail(ctx context.Context, in model.APIContentGetDetailInput) (out *model.APIContentGetDetailOutput, err error)
}

var localAPIIArticle APIIArticle

func APIArticle() APIIArticle {
	if localAPIIArticle == nil {
		panic("implement not found for interface APIArticle, forgot register?")
	}
	return localAPIIArticle
}

func RegisterAPIIArticle(i APIIArticle) {
	localAPIIArticle = i
}