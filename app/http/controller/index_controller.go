package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/klein/go-mvc/pkg/session"
)

type IndexController struct {
	BaseController
}


func (i *IndexController) Index(ctx *gin.Context) {

	session.Set("session", "你好")
	data := gin.H{
		"message": "welcome",
		"session": session.Get("session"),
	}
	i.SendJsonResponse(ctx, 203, "success", data)
}