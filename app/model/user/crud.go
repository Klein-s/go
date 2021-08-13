package user

import (
	"github.com/klein/go-mvc/pkg/logger"
	"github.com/klein/go-mvc/pkg/model"
	"github.com/klein/go-mvc/pkg/types"
)

func (user *User) Create() (err error)  {
	if err = model.DB.Create(&user).Error ; err != nil {
		logger.LogError(err)
		return err
	}
	return nil
}

//根据 id 获取用户信息
func Get(idstr string) (User, error)  {
	var user User
	id := types.StringToInt(idstr)
	if err := model.DB.First(&user, id).Error; err != nil {
		return user, err
	}

	return user, nil

}

//根据邮箱获取用户信息
func GetByEmail(email string) (User, error)  {
	var user User
	if err := model.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}
