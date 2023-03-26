package logic

//logic 其实就是存放业务层的代码

import (
	"go_server/dao/mysql"
	"go_server/models"
	"go_server/pkg/jwt"
	"go_server/pkg/snowflake"
)

func Login(login *models.ParamLogin) (string, error) {
	user := models.User{
		Email:    login.Email,
		Password: login.Password,
	}
	if err := mysql.Login(&user); err != nil {
		return "", err
	}
	return jwt.GenToken(user.Email, user.UserId)
}

func Register(register *models.ParamRegister) (err error) {
	// 判断用户是否存在
	err = mysql.CheckUserExist(register.Email)
	if err != nil {
		// db 出错
		return err
	}
	// 生成userid
	userId := snowflake.GenId()
	// 构造一个User db对象
	user := models.User{
		UserId:   userId,
		Username: register.UserName,
		Password: register.Password,
		Email:    register.Email,
	}
	// 保存数据库
	err = mysql.InsertUser(&user)
	if err != nil {
		return err
	}
	return
}
