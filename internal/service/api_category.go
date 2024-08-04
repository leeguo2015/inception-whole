// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package service

import (
	"context"
	"inception-whole/internal/model"
	// "inception-whole/internal/model/entity"
)

type APIICategory interface {
	GetALL(ctx context.Context, Type string)  (list []*model.APICategoryItem, err error)
}

var localAPIICategory APIICategory

func APICategory() APIICategory {
	if localCategory == nil {
		panic("implement not found for interface ICategory, forgot register?")
	}
	return localAPIICategory
}

func RegisterAPICCategory(i APIICategory) {
	localAPIICategory = i
}
