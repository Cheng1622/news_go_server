package mysql

import (
	"database/sql"
	"go_server/models"

	"go.uber.org/zap"
)

func GetListList() (list []*models.List, err error) {
	sqlStr := "select list_id,content,create_time from list"
	err = db.Select(&list, sqlStr)
	if err != nil {
		// 空数据的时候 不算错误 只是没有板块而已
		if err == sql.ErrNoRows {
			zap.L().Warn("no list ")
			err = nil
		}
	}
	return
}
func GetListListLast() (list []*models.List, err error) {
	sqlStr := "select list_id,content,create_time" +
		" from list order by id DESC limit 1"
	err = db.Select(&list, sqlStr)
	if err != nil {
		// 空数据的时候 不算错误 只是没有板块而已
		if err == sql.ErrNoRows {
			zap.L().Warn("no list ")
			err = nil
		}
	}
	return

}

func GetListById(id int64) (list *models.List, err error) {
	list = new(models.List)
	sqlStr := "select list_id,content,create_time " +
		"from list where list_id=?"
	err = db.Get(list, sqlStr, id)
	if err != nil {
		// 空数据的时候 不算错误 只是没有板块而已
		if err == sql.ErrNoRows {
			zap.L().Warn("no list ")
			err = nil
		}
	}
	return list, err

}

func InsertList(list *models.List) error {
	sqlstr := `insert into list(list_id,content) values(?,?)`
	_, err := db.Exec(sqlstr, list.Id, list.Content)

	if err != nil {
		zap.L().Error("InsertPost dn error", zap.Error(err))
		return err
	}
	return nil
}
