package service

import (
	"github.com/Cheng1622/news_go_server/internal/model"
	"github.com/Cheng1622/news_go_server/internal/model/requ"
	"github.com/Cheng1622/news_go_server/pkg/mysql"
)

// ExampleService
type DictService interface {
	PostDict(dict *model.Dict) error                                                  // 创建
	GetDictList(pageInfo *requ.PageList) (data []*model.Dict, total int64, err error) // 列表
	DeleteDict(id *requ.DictId) error                                                 // 删除
	DeleteDictAll(dictIds []uint) error                                               // 批量删除
	PutDict(dict *model.Dict) error                                                   // 更新
}

type Dict struct{}

// NewDictService 构造函数
func NewDictService() DictService {
	return Dict{}
}

// PostDict 创建
func (dt Dict) PostDict(dict *model.Dict) error {
	err := mysql.DB.Create(&dict).Error
	return err
}

// GetDictList 列表
func (dt Dict) GetDictList(pageInfo *requ.PageList) (data []*model.Dict, total int64, err error) {
	// gorm 获获列表数据
	limit := pageInfo.Size
	offset := pageInfo.Size * (pageInfo.Page - 1)
	// 创建db
	db := mysql.DB.Debug().Order("id desc").Model(&model.Dict{})
	// 如果有条件搜索 下方会自动创建搜索语句
	if pageInfo.ID != 0 {
		db = db.Where("`id` = ?", pageInfo.ID)
	}
	// name
	if pageInfo.Name != "" {
		db = db.Where("`name` LIKE ?", "%"+pageInfo.Name+"%")
	}
	// keyword
	if pageInfo.KeyWord != "" {
		db = db.Where("`key_word` = ?", pageInfo.KeyWord)
	}
	// desc
	if pageInfo.Desc != "" {
		db = db.Where("`desc` LIKE ?", "%"+pageInfo.Desc+"%")
	}
	err = db.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}
	err = db.Limit(limit).Offset(offset).Find(&data).Error

	return data, total, err
}

// DeleteDict 删除
func (dt Dict) DeleteDict(id *requ.DictId) error {
	// 软删除,根据主键软删除
	err := mysql.DB.Delete(&model.Dict{}, id).Error
	return err
}

// DeleteDictAll 批量删除
func (dt Dict) DeleteDictAll(dictIds []uint) (err error) {
	// 软删除
	err = mysql.DB.Where("id IN (?)", dictIds).Delete(&model.Dict{}).Error
	return err
}

// PutDict 更新
func (dt Dict) PutDict(dict *model.Dict) error {
	// 根据id更新
	err := mysql.DB.Model(dict).Where("id = ?", dict.ID).Updates(&dict).Error
	return err
}
