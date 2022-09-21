package router

import (
	"web_stock/controller"

	"github.com/gin-gonic/gin"
)

func Router(e *gin.Engine) {
	e.LoadHTMLGlob("templates/html/*.html")
	e.Static("/static", "templates/html")
	e.GET("/register", controller.Register)
	e.GET("/login", controller.Login)
	e.GET("/data", controller.Data)
}
