package mysql

import (
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	"errors"
	"go_server/models"

	"go.uber.org/zap"
)

const serect = "cjic.ga"

// 定义 error的常量方便判断
var (
	UserAleadyExists = errors.New("用户已存在")
	WrongPassword    = errors.New("密码不正确")
	UserNoExists     = errors.New("用户不存在")
)

func GetEmailById(id int64) (email string, err error) {
	sqlStr := "select email " +
		" from user where user_id=?"
	err = db.Get(&email, sqlStr, id)
	return email, err
}
func GetUserNewsById(id int64) (isnews int32, err error) {
	sqlStr := "select isnews " +
		" from post where post_id=?"
	err = db.Get(&isnews, sqlStr, id)
	return isnews, err
}

// dao层 其实就是将数据库操作 封装为函数 等待logic层 去调用她

func InsertUser(user *models.User) error {
	// 密码要加密保存
	user.Password = encryptPassword(user.Password)
	sqlstr := `insert into user(user_id,username,password,email) values(?,?,?,?)`
	_, err := db.Exec(sqlstr, user.UserId, user.Username, user.Password, user.Email)
	if err != nil {
		zap.L().Error("InsertUser dn error", zap.Error(err))
		return err
	}
	return nil
}

func Login(user *models.User) error {
	oldPassword := user.Password
	sqlStr := `select user_id,username,password,email from user where email=?`
	err := db.Get(user, sqlStr, user.Email)
	if err == sql.ErrNoRows {
		return UserNoExists
	}
	if err != nil {
		return err
	}
	if encryptPassword(oldPassword) != user.Password {
		return WrongPassword
	}
	return nil
}
func GetUserByEmail(user *models.User) (a *models.User, err error) {
	sqlStr := `select user_id,username,password,email from user where email=?`
	err = db.Get(user, sqlStr, user.Email)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// CheckUserExist 检查数据库是否有该用户名
func CheckUserExist(email string) error {
	sqlstr := `select count(user_id) from user where email = ?`
	var count int
	err := db.Get(&count, sqlstr, email)
	if err != nil {
		zap.L().Error("CheckUserExist dn error", zap.Error(err))
		return err
	}
	if count > 0 {
		return UserAleadyExists
	}
	return nil
}

// 加密密码
func encryptPassword(password string) string {
	h := md5.New()
	h.Write([]byte(serect))
	return hex.EncodeToString(h.Sum([]byte(password)))
}
