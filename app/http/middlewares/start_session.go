package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/klein/go-mvc/pkg/session"
)

//startSession 开启session 会话控制
func StartSession() gin.HandlerFunc  {
	return func(c *gin.Context) {

		//启用会话
		session.StartSession(c)

		// 继续处理请求
		c.Next()
	}
}