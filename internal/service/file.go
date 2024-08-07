// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package service

import (
	"context"
	"inception-whole/internal/model"
)

type IFile interface {
	Upload(ctx context.Context, in model.FileUploadInput) (*model.FileUploadOutput, error)
}

var localFile IFile

func File() IFile {
	if localFile == nil {
		panic("implement not found for interface IFile, forgot register?")
	}
	return localFile
}

func RegisterFile(i IFile) {
	localFile = i
}
