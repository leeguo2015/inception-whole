// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package service

import (
	"context"
	"inception-whole/internal/model"
)

type IReply interface {
	Create(ctx context.Context, in model.ReplyCreateInput) error
	Delete(ctx context.Context, id uint) error
	DeleteByUserContentId(ctx context.Context, userId, contentId uint) error
	GetList(ctx context.Context, in model.ReplyGetListInput) (out *model.ReplyGetListOutput, err error)
}

var localReply IReply

func Reply() IReply {
	if localReply == nil {
		panic("implement not found for interface IReply, forgot register?")
	}
	return localReply
}

func RegisterReply(i IReply) {
	localReply = i
}
