package user

import (
	"github.com/klein/go-mvc/app/model"
	"github.com/klein/go-mvc/pkg/password"
)

//用户模型
type User struct {
	model.BaseModel

	Name string `gorm:"type:varchar(255);not null;unique" valid:"name"`
	Email string `gorm:"type:varchar(255);default:NULL;unique" valid:"email"`
	Password string `gorm:"type:varchar(255)" valid:"password"`

	PasswordConfirm string `gorm:"-" valid:"password_confirm"`
	Role string `gorm:"type:varchar(255); default:NULL"`
}

func (u User) HasRole(role string) bool  {
	if u.Role == role{
		return true
	}
	return false
}

// ComparePassword 对比密码是否匹配
func (u User) ComparePassword(_password string) bool  {
	return password.CheckHas(_password, u.Password)
}
