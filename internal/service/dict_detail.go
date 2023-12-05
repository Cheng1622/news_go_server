package service

import (
	"github.com/Cheng1622/news_go_server/internal/model"
	"github.com/Cheng1622/news_go_server/internal/model/requ"
	"github.com/Cheng1622/news_go_server/pkg/mysql"
)

// DictDetailsService
type DictDetailsService interface {
	PostDictDetails(dictDetail *model.DictDetail) error                                                  // 创建
	GetDictDetailsList(pageInfo *requ.DictDetailList) (data []*model.DictDetail, total int64, err error) // 列表
	DeleteDictDetails(ID *requ.DictId) error                                                             // 删除
	DeleteDictDetailsAll(dictDetailsIds []uint) error                                                    // 批量删除
	PutDictDetails(dictDetail *model.DictDetail) error                                                   // 更新
}

type DictDetails struct{}

// NewDictDetailsService 构造函数
func NewDictDetailsService() DictDetailsService {
	return DictDetails{}
}

// PostDictDetails 创建
func (ds DictDetails) PostDictDetails(dictDetail *model.DictDetail) error {
	err := mysql.DB.Create(&dictDetail).Error
	return err
}

// GetDictDetailsList 列表
func (ds DictDetails) GetDictDetailsList(pageInfo *requ.DictDetailList) (data []*model.DictDetail, total int64, err error) {
	// gorm 获获列表数据
	limit := pageInfo.Size
	offset := pageInfo.Size * (pageInfo.Page - 1)
	KeyWord := pageInfo.KeyWord

	// sql语句
	sql := "SELECT * FROM dict_detail WHERE dict_id IN ( " +
		"SELECT id FROM dict WHERE key_word = ? and status IS True  ) and deleted_at IS NULL order by sort limit ?,?;"

	if err = mysql.DB.Raw(sql, KeyWord, offset, limit).Scan(&data).Error; err != nil {
		return data, 0, err
	}

	db := mysql.DB.Order("sort asc").Model(&model.DictDetail{})

	err = db.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	return data, total, err
}

// DeleteDictDetails 删除
func (ds DictDetails) DeleteDictDetails(ID *requ.DictId) error {
	// 软删除,根据主键软删除
	err := mysql.DB.Delete(&model.DictDetail{}, ID).Error
	return err
}

// DeleteExampleAll 批量删除
func (ds DictDetails) DeleteDictDetailsAll(dictDetailsIds []uint) (err error) {
	// 软删除
	err = mysql.DB.Where("id IN (?)", dictDetailsIds).Delete(&model.DictDetail{}).Error
	return err
}

// PutExample 更新
func (ds DictDetails) PutDictDetails(dictDetail *model.DictDetail) error {
	// 根据id更新
	err := mysql.DB.Model(dictDetail).Where("id = ?", dictDetail.ID).Updates(&dictDetail).Error
	return err
}
