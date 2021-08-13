package http

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/klein/go-mvc/app/http/middlewares"
	"github.com/klein/go-mvc/pkg/config"
	"github.com/klein/go-mvc/pkg/session"
)

/**
中间件
 */
func Middleware(r *gin.Engine)  {

	//初始化session
	r.Use(sessions.Sessions(config.GetString("app.key"), session.Store()))
	r.Use(middlewares.StartSession())

}
