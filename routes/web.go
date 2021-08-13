package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/klein/go-mvc/app/http"
	"github.com/klein/go-mvc/app/http/controller"
)

func RegisterWebRoutes() *gin.Engine {
	r := gin.Default()

	//初始化中间件
	http.Middleware(r)

	ic := new(controller.IndexController)
	r.GET("/", ic.Index)

	return r
}