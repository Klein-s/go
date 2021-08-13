package controller

import "github.com/gin-gonic/gin"

type BaseController struct {

}
type Response struct {
	Code int `json:"code"`
	Message string `json:"message"`
	Data interface{} `json:"data"`
}

func (*BaseController) SendJsonResponse(c *gin.Context, code int, message string, data interface{})  {
	c.JSON(code, Response{
		Code: code,
		Message: message,
		Data: data,
	})
}