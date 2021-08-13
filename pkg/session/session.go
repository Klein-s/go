package session

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)


// session 当前会话

var Session sessions.Session

// startSession 初始化会话。 中间件当中调用
func StartSession(c *gin.Context)  {

	//初始化session
	Session = sessions.Default(c)
}

//set 写入键值对应的会话数据
func Set(key string, value interface{}) {
	Session.Set(key, value)
	_ = Session.Save()
}

//get 获取会话数据， 获取数据时请做类型检测

func Get(key string) interface{}  {
	return Session.Get(key)
}

// forget 删除某个会话项
func Forget(key string)  {
	Session.Delete(key)
}

// flush 删除当前会话
func Flush() {
	Session.Clear()
}