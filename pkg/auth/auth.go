package auth

import (
	"errors"
	"github.com/klein/go-mvc/app/model/user"
	"github.com/klein/go-mvc/pkg/session"
	"gorm.io/gorm"
)

func _getUID() string  {
	_uid := session.Get("uid")

	uid, ok := _uid.(string)
	if ok && len(uid) > 0 {
		return uid
	}
	return ""
}

//获取登录用户信息
func User() user.User  {
	uid := _getUID()
	if len(uid) > 0 {
		_user, err := user.Get(uid)
		if err == nil {
			return _user
		}
	}
	return user.User{}
}

//Attempt 尝试登录

func Attempt(email string, password string) error  {
	//根据Email 获取用户
	_user, err := user.GetByEmail(email)
	//出现错误

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return errors.New("账号不存在或密码错误")
		} else {
			return errors.New("内部错误，请稍后尝试")
		}
	}

	// 匹配密码
	if !_user.ComparePassword(password) {
		return errors.New("账号不存在或密码错误")
	}

	//用户登录 保存会话
	session.Set("uid", _user.GetStringID())

	return nil
}

//登录指定用户
func Login(_user user.User)  {
	session.Set("uid", _user.GetStringID())
}

// logout 退出登录
func Logout()  {
	session.Forget("uid")
}

//check 检查是否登录
func Check() bool  {
	return len(_getUID()) > 0
}