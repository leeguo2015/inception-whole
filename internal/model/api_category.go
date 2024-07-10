package model

type APICategoryItem struct {
	Id      uint   `json:"id"`      // 分类ID，自增主键
	Name    string `json:"name"`    // 分类名称
	Content string `json:"content"` // 详细介绍
	Brief   string `json:"brief"`   // 简述

}
