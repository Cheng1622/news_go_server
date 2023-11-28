package model

import "gorm.io/gorm"

// 字典详情表结构体
type DictDetail struct {
	gorm.Model
	Label  string `gorm:"type:varchar(50);comment:'展示值'" json:"label"`
	Value  int    `gorm:"type:varchar(50);comment:'字典值'" json:"value"`
	Sort   int    `gorm:"type:varchar(50);comment:'排序'" json:"sort"`
	Desc   string `gorm:"type:varchar(100);comment:'说明'" json:"desc"`
	DictID uint   `json:"dict_id" gorm:"comment:字典分类id;"`
}
