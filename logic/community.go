package logic

import (
	"go_server/dao/mysql"
	"go_server/models"
)

func GetCommunityList() (communityList []*models.Community, err error) {
	return mysql.GetCommunityList()

}

func GetCommunityById(id int64) (model *models.Community, err error) {
	return mysql.GetCommunityById(id)

}
