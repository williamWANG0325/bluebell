package logic

import (
	"bluebell/dao/mysql"
	"bluebell/models"
	"bluebell/pkg/jwt"
	"bluebell/pkg/snowflake"
)

func SignUp(p *models.ParamSignUp) (err error) {
	// 判断用户是否存在
	if err := mysql.CheckUserExist(p.Username); err != nil {
		return err
	}
	// 构造用户实例
	user := &models.User{
		UserID:   snowflake.GenID(),
		Username: p.Username,
		Password: p.Password,
	}
	// 保存进数据库
	err = mysql.InsertUser(user)
	return
}

func Login(p *models.ParamLogin) (user *models.User, err error) {
	user = &models.User{
		Username: p.Username,
		Password: p.Password,
	}
	// 判断用户是否存在
	if err := mysql.Login(user); err != nil {
		return nil, err
	}
	user.Token, err = jwt.GenToken(user.UserID, user.Username)
	return
}
