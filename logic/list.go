package logic

import (
	"go_server/dao/mysql"
	"go_server/models"
)

func GetListList() (communityList []*models.List, err error) {
	return mysql.GetListList()

}
func GetListListLast() (communityList []*models.List, err error) {
	return mysql.GetListListLast()

}

func GetListById(id int64) (model *models.List, err error) {
	return mysql.GetListById(id)

}
